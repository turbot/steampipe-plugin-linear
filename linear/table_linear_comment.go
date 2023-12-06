package linear

import (
	"context"

	"github.com/turbot/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableLinearComment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_comment",
		Description: "Linear Comment",
		List: &plugin.ListConfig{
			Hydrate: listComments,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "created_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "updated_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:    "body",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getComment,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier of the entity.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "archived_at",
				Description: "The time at which the entity was archived. Null if the entity has not been archived.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "body",
				Description: "The comment content in markdown format.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "body_data",
				Description: "The comment content as a Prosemirror document.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time at which the entity was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "edited_at",
				Description: "The time user edited the comment.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "reaction_data",
				Description: "Emoji reaction summary, grouped by emoji type.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "updated_at",
				Description: "The last time at which the entity was meaningfully updated, i.e., for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "url",
				Description: "Comment's URL.",
				Type:        proto.ColumnType_STRING,
			},
			// user is a keyword, so here transform function has been used
			{
				Name:        "comment_user",
				Description: "The user who wrote the comment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("User"),
			},
			{
				Name:        "parent",
				Type:        proto.ColumnType_JSON,
				Description: "The parent comment under which the current comment is nested.",
			},
			{
				Name:        "issue",
				Type:        proto.ColumnType_JSON,
				Description: "The issue that the comment is associated with.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The comment's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
		},
	}
}

// LIST FUNCTION

func listComments(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_comment.listComments", "connection_error", err)
		return nil, err
	}
	var endCursor string

	// set page size
	var pageSize int = int(conn.pageSize)
	if d.QueryContext.Limit != nil {
		if int(*d.QueryContext.Limit) < pageSize {
			pageSize = int(*d.QueryContext.Limit)
		}
	}

	// set the requested filters
	filters := setCommentFilters(d, ctx)

	for {
		listCommentResponse, err := gql.ListComments(ctx, conn.client, pageSize, endCursor, true, &filters)
		if err != nil {
			plugin.Logger(ctx).Error("linear_comment.listComments", "api_error", err)
			return nil, err
		}
		for _, node := range listCommentResponse.Comments.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listCommentResponse.Comments.PageInfo.HasNextPage {
			break
		}
		endCursor = *listCommentResponse.Comments.PageInfo.EndCursor
	}

	return nil, nil
}

// HYDRATE FUNCTION

func getComment(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_comment.getComment", "connection_error", err)
		return nil, err
	}

	getCommentResponse, err := gql.GetComment(ctx, conn.client, &id)
	if err != nil {
		plugin.Logger(ctx).Error("linear_comment.getComment", "api_error", err)
		return nil, err
	}

	return getCommentResponse.Comment, nil
}

// Set the requested filter
func setCommentFilters(d *plugin.QueryData, ctx context.Context) gql.CommentFilter {
	var filter gql.CommentFilter
	if d.Quals["created_at"] != nil {
		createdAt := &gql.DateComparator{}
		for _, q := range d.Quals["created_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				createdAt.Eq = timestamp
			case ">":
				createdAt.Gt = timestamp
			case ">=":
				createdAt.Gte = timestamp
			case "<":
				createdAt.Lt = timestamp
			case "<=":
				createdAt.Lte = timestamp
			}
		}
		filter.CreatedAt = createdAt
	}
	if d.Quals["updated_at"] != nil {
		updatedAt := &gql.DateComparator{}
		for _, q := range d.Quals["updated_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				updatedAt.Eq = timestamp
			case ">":
				updatedAt.Gt = timestamp
			case ">=":
				updatedAt.Gte = timestamp
			case "<":
				updatedAt.Lt = timestamp
			case "<=":
				updatedAt.Lte = timestamp
			}
		}
		filter.UpdatedAt = updatedAt
	}
	if d.EqualsQuals["body"] != nil {
		body := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("body")),
		}
		filter.Body = body
	}
	return filter
}
