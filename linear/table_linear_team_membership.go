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
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTeamMembership,
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
			// user is a keyword, so here transform function has been used
			{
				Name:        "membership_user",
				Type:        proto.ColumnType_JSON,
				Description: "The user that the membership is associated with.",
				Transform:   transform.FromField("User"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The issue label's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
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

	// set page size
	var pageSize int = int(conn.pageSize)
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
		case "membership_user":
			includeUser = false
		}
	}

	for {
		listTeamMembershipResponse, err := gql.ListTeamMemberships(ctx, conn.client, pageSize, endCursor, true, &includeTeam, &includeUser)
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

func getTeamMembership(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeTeam, includeUser := true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "team":
			includeTeam = false
		case "membership_user":
			includeUser = false
		}
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_team_membership.getTeamMembership", "connection_error", err)
		return nil, err
	}

	getTeamMembershipResponse, err := gql.GetTeamMembership(ctx, conn.client, &id, &includeTeam, &includeUser)
	if err != nil {
		plugin.Logger(ctx).Error("linear_team_membership.getTeamMembership", "api_error", err)
		return nil, err
	}

	return getTeamMembershipResponse.TeamMembership, nil
}
