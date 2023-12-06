package linear

import (
	"context"

	"github.com/turbot/steampipe-plugin-linear/gql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableLinearIntegration(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_integration",
		Description: "Linear Integration",
		List: &plugin.ListConfig{
			Hydrate: listIntegrations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getIntegration,
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
				Name:        "service",
				Description: "The integration's type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "The last time at which the entity was meaningfully updated, i.e., for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "team",
				Description: "The team that the integration is associated with.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "creator",
				Description: "The user that added the integration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "organization",
				Description: "The organization that the integration is associated with.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The integration's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
		},
	}
}

// LIST FUNCTION

func listIntegrations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_integration.listIntegrations", "connection_error", err)
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

	for {
		listIntegrationResponse, err := gql.ListIntegrations(ctx, conn.client, pageSize, endCursor, true)
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

// HYDRATE FUNCTION

func getIntegration(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_integration.getIntegration", "connection_error", err)
		return nil, err
	}

	getIntegrationResponse, err := gql.GetIntegration(ctx, conn.client, &id)
	if err != nil {
		plugin.Logger(ctx).Error("linear_integration.getIntegration", "api_error", err)
		return nil, err
	}

	return getIntegrationResponse.Integration, nil
}
