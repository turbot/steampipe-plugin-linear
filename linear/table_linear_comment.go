package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinearComment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_comment",
		Description: "Linear Comment",
		List: &plugin.ListConfig{
			Hydrate: listComments,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "id",
					Require: plugin.Optional,
				},
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
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the entity.",
			},
			{
				Name:        "archived_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was archived. Null if the entity has not been archived.",
			},
			{
				Name:        "body",
				Type:        proto.ColumnType_STRING,
				Description: "The comment content in markdown format.",
			},
			{
				Name:        "body_data",
				Type:        proto.ColumnType_STRING,
				Description: "The comment content as a Prosemirror document.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was created.",
			},
			{
				Name:        "edited_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time user edited the comment.",
			},
			{
				Name:        "reaction_data",
				Type:        proto.ColumnType_JSON,
				Description: "Emoji reaction summary, grouped by emoji type.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e., for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "Comment's URL.",
			},
			{
				Name:        "user",
				Type:        proto.ColumnType_JSON,
				Description: "The user who wrote the comment.",
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
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

func listComments(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_comment.listComments", "connection_error", err)
		return nil, err
	}
	var endCursor string
	var pageSize int = 50
	if d.QueryContext.Limit != nil {
		if int(*d.QueryContext.Limit) < pageSize {
			pageSize = int(*d.QueryContext.Limit)
		}
	}

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeIssue, includeParent, includeUser := true, true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "issue":
			includeIssue = false
		case "parent":
			includeParent = false
		case "user":
			includeUser = false
		}
	}

	// set the requested filters
	filters := setCommentFilters(d, ctx)

	for {
		listCommentResponse, err := gql.ListComment(ctx, conn, pageSize, endCursor, true, &filters, &includeIssue, &includeParent, &includeUser)
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

// Set the requested filter
func setCommentFilters(d *plugin.QueryData, ctx context.Context) gql.CommentFilter {
	var filter gql.CommentFilter
	if d.EqualsQuals["id"] != nil {
		id := &gql.IDComparator{
			Eq: types.String(d.EqualsQualString("id")),
		}
		filter.Id = id
	}
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
