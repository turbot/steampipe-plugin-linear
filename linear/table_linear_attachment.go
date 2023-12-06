package linear

import (
	"context"

	"github.com/turbot/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableLinearAttachment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_attachment",
		Description: "Linear Attachment",
		List: &plugin.ListConfig{
			Hydrate: listAttachments,
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
					Name:    "title",
					Require: plugin.Optional,
				},
				{
					Name:    "subtitle",
					Require: plugin.Optional,
				},
				{
					Name:    "source_type",
					Require: plugin.Optional,
				},
				{
					Name:    "url",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAttachment,
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
				Name:        "created_at",
				Description: "The time at which the entity was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "group_by_source",
				Description: "Indicates if attachments for the same source application should be grouped in the Linear UI.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "metadata",
				Description: "Custom metadata related to the attachment.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "source",
				Description: "Information about the source which created the attachment.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "source_type",
				Description: "An accessor helper to source.type, defines the source type of the attachment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "subtitle",
				Description: "Content for the subtitle line in the Linear attachment widget.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "The last time at which the entity was meaningfully updated, i.e., for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "url",
				Description: "Location of the attachment which is also used as an identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creator",
				Description: "The creator of the attachment.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "issue",
				Description: "The issue this attachment belongs to.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The attachment's title.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

// LIST FUNCTION

func listAttachments(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_attachment.listAttachments", "connection_error", err)
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
	filters := setAttachmentFilters(d, ctx)

	for {
		listAttachmentResponse, err := gql.ListAttachments(ctx, conn.client, pageSize, endCursor, true, &filters)
		if err != nil {
			plugin.Logger(ctx).Error("linear_attachment.listAttachments", "api_error", err)
			return nil, err
		}
		for _, node := range listAttachmentResponse.Attachments.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listAttachmentResponse.Attachments.PageInfo.HasNextPage {
			break
		}
		endCursor = *listAttachmentResponse.Attachments.PageInfo.EndCursor
	}

	return nil, nil
}

// HYDRATE FUNCTION

func getAttachment(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_attachment.getAttachment", "connection_error", err)
		return nil, err
	}

	getAttachmentResponse, err := gql.GetAttachment(ctx, conn.client, &id)
	if err != nil {
		plugin.Logger(ctx).Error("linear_attachment.getAttachment", "api_error", err)
		return nil, err
	}

	return getAttachmentResponse.Attachment, nil
}

// Set the requested filter
func setAttachmentFilters(d *plugin.QueryData, ctx context.Context) gql.AttachmentFilter {
	var filter gql.AttachmentFilter
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
	if d.EqualsQuals["title"] != nil {
		title := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("title")),
		}
		filter.Title = title
	}
	if d.EqualsQuals["subtitle"] != nil {
		subtitle := &gql.NullableStringComparator{
			Eq: types.String(d.EqualsQualString("subtitle")),
		}
		filter.Subtitle = subtitle
	}
	if d.EqualsQuals["source_type"] != nil {
		source_type := &gql.SourceTypeComparator{
			Eq: types.String(d.EqualsQualString("source_type")),
		}
		filter.SourceType = source_type
	}
	if d.EqualsQuals["url"] != nil {
		url := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("url")),
		}
		filter.Url = url
	}

	return filter
}
