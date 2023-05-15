package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinearIntegration(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_integration",
		Description: "Linear Integration",
		List: &plugin.ListConfig{
			Hydrate: listIntegrations,
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
				Name:        "service",
				Type:        proto.ColumnType_STRING,
				Description: "The integration's type.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e., for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "team",
				Type:        proto.ColumnType_JSON,
				Description: "The team that the integration is associated with.",
			},
			{
				Name:        "creator",
				Type:        proto.ColumnType_JSON,
				Description: "The user that added the integration.",
			},
			{
				Name:        "organization",
				Type:        proto.ColumnType_JSON,
				Description: "The organization that the integration is associated with.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The integration's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

func listIntegrations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_integration.listIntegrations", "connection_error", err)
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
	includeCreator, includeOrganization, includeTeam := true, true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "creator":
			includeCreator = false
		case "organization":
			includeOrganization = false
		case "team":
			includeTeam = false
		}
	}

	for {
		listIntegrationResponse, err := gql.ListIntegration(ctx, conn, pageSize, endCursor, true, &includeCreator, &includeOrganization, &includeTeam)
		if err != nil {
			plugin.Logger(ctx).Error("linear_integration.listIntegrations", "api_error", err)
			return nil, err
		}
		for _, node := range listIntegrationResponse.Integrations.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listIntegrationResponse.Integrations.PageInfo.HasNextPage {
			break
		}
		endCursor = *listIntegrationResponse.Integrations.PageInfo.EndCursor
	}

	return nil, nil
}
