package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinearTeamMembership(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_team_membership",
		Description: "Linear Team Membership",
		List: &plugin.ListConfig{
			Hydrate: listTeamMemberships,
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
				Name:        "owner",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the user is the owner of the team.",
			},
			{
				Name:        "sort_order",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The order of the item in the user's team list.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e., for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "team",
				Type:        proto.ColumnType_JSON,
				Description: "The team that the membership is associated with.",
			},
			{
				Name:        "user",
				Type:        proto.ColumnType_JSON,
				Description: "The user that the membership is associated with.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The issue label's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

func listTeamMemberships(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_team_membership.listTeamMemberships", "connection_error", err)
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
	includeTeam, includeUser := true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "team":
			includeTeam = false
		case "user":
			includeUser = false
		}
	}

	for {
		listTeamMembershipResponse, err := gql.ListTeamMembership(ctx, conn, pageSize, endCursor, true, &includeTeam, &includeUser)
		if err != nil {
			plugin.Logger(ctx).Error("linear_team_membership.listTeamMemberships", "api_error", err)
			return nil, err
		}

		for _, node := range listTeamMembershipResponse.TeamMemberships.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listTeamMembershipResponse.TeamMemberships.PageInfo.HasNextPage {
			break
		}
		endCursor = *listTeamMembershipResponse.TeamMemberships.PageInfo.EndCursor
	}

	return nil, nil
}
