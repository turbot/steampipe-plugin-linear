package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLinearAttachment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_attachment",
		Description: "Linear Attachment",
		List: &plugin.ListConfig{
			Hydrate: listAttachments,
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
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was created.",
			},
			{
				Name:        "group_by_source",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates if attachments for the same source application should be grouped in the Linear UI.",
			},
			{
				Name:        "metadata",
				Type:        proto.ColumnType_JSON,
				Description: "Custom metadata related to the attachment.",
			},
			{
				Name:        "source",
				Type:        proto.ColumnType_JSON,
				Description: "Information about the source which created the attachment.",
			},
			{
				Name:        "source_type",
				Type:        proto.ColumnType_STRING,
				Description: "An accessor helper to source.type, defines the source type of the attachment.",
			},
			{
				Name:        "subtitle",
				Type:        proto.ColumnType_STRING,
				Description: "Content for the subtitle line in the Linear attachment widget.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e., for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "Location of the attachment which is also used as an identifier.",
			},
			{
				Name:        "creator",
				Type:        proto.ColumnType_JSON,
				Description: "The creator of the attachment.",
			},
			{
				Name:        "issue",
				Type:        proto.ColumnType_JSON,
				Description: "The issue this attachment belongs to.",
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

func listAttachments(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_attachment.listAttachments", "connection_error", err)
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
	includeCreator, includeIssue := true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "creator":
			includeCreator = false
		case "issue":
			includeIssue = false
		}
	}

	// set the requested filters
	filters := setAttachmentFilters(d, ctx)

	for {
		listAttachmentResponse, err := gql.ListAttachment(ctx, conn, pageSize, endCursor, true, &filters, &includeCreator, &includeIssue)
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

// Set the requested filter
func setAttachmentFilters(d *plugin.QueryData, ctx context.Context) gql.AttachmentFilter {
	var filter gql.AttachmentFilter
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
