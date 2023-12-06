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

func tableLinearTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_team",
		Description: "Linear Team",
		List: &plugin.ListConfig{
			Hydrate: listTeams,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "key",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
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
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTeam,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "A unique identifier for the team.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "archived_at",
				Description: "The time at which the team was archived (if it has been).",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "auto_archive_period",
				Description: "The period (in months) after which automatically closed and completed issues are automatically archived.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "auto_close_period",
				Description: "The period (in months) after which issues are automatically closed. Null/undefined means this feature is disabled.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "auto_close_state_id",
				Description: "The canceled workflow state which auto closed issues will be set to. Defaults to the first canceled state.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "color",
				Description: "The team's color.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time at which the team was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "cycle_calender_url",
				Description: "Calendar feed URL (iCal) for cycles.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cycle_cooldown_time",
				Description: "The cooldown time after each cycle in weeks.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "cycle_duration",
				Description: "The duration of a cycle in weeks.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "cycle_issue_auto_assign_completed",
				Description: "Whether completed issues are automatically assigned to the current cycle.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "cycle_issue_auto_assign_started",
				Description: "Whether started issues are automatically assigned to the current cycle.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "cycle_lock_to_active",
				Description: "Whether only issues with cycles in Active Issues are allowed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "cycle_start_day",
				Description: "The day of the week that a new cycle starts.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cycles_enabled",
				Description: "Whether the team uses cycles.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "default_issue_estimate",
				Description: "The default estimate for unestimated issues.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "default_template_for_members_id",
				Description: "The ID of the default template to use for new issues created by members of the team.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "default_template_for_non_members_id",
				Description: "The ID of the default template to use for new issues created by non-members of the team.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The team's description.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "group_issue_history",
				Description: "Whether to group recent issue history entries.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "icon",
				Description: "The team's icon.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "invite_hash",
				Description: "A unique hash for the team to be used in invite URLs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "issue_estimation_allow_zero",
				Description: "Whether to allow zeros in issues estimates.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "issue_estimation_extended",
				Description: "Whether to add additional points to the estimate scale.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "issue_estimation_type",
				Description: "The issue estimation type to use.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "issue_ordering_no_priority_first",
				Description: "Whether issues without priority should be sorted first.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "issue_sort_order_default_to_bottom",
				Description: "Whether to move issues to bottom of the column when changing state.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "key",
				Description: "The team's unique key. The key is used in URLs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The team's name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "private",
				Description: "Whether the team is private or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "require_priority_to_leave_triage",
				Description: "Whether an issue needs to have a priority set before leaving triage.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "slack_issue_comments",
				Description: "Whether to send new issue comment notifications to Slack.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "slack_issue_statuses",
				Description: "Whether to send new issue status updates to Slack.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "slack_new_issue",
				Description: "Whether to send new issue notifications to Slack.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "timezone",
				Description: "The timezone of the team. Defaults to 'America/Los_Angeles'",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "triage_enabled",
				Description: "Whether triage mode is enabled for the team or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "upcoming_cycle_count",
				Description: "How many upcoming cycles to create.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "updated_at",
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "active_cycle",
				Description: "Team's currently active cycle.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "default_issue_state",
				Description: "The default workflow state into which issues are set when they are opened by team members.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "default_template_for_members",
				Description: "The default template to use for new issues created by members of the team.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "default_template_for_non_members",
				Description: "The default template to use for new issues created by non-members of the team.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "draft_workflow_state",
				Description: "The workflow state into which issues are moved when a PR has been opened as draft.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "integrations_settings",
				Description: "Settings for all integrations associated with that team.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "marked_as_duplicate_workflow_state",
				Description: "The workflow state into which issues are moved when they are marked as a duplicate of another issue. Defaults to the first canceled state.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "organization",
				Description: "The organization that the team is associated with.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "review_workflow_state",
				Description: "The workflow state into which issues are moved when a review has been requested for the PR.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "start_workflow_state",
				Description: "The workflow state into which issues are moved when a PR has been opened.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "triage_issue_state",
				Description: "The workflow state into which issues are set when they are opened by non-team members or integrations if triage is enabled.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The team's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

// LIST FUNCTION

func listTeams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_team.listTeams", "connection_error", err)
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
	filters := setTeamFilters(d, ctx)

	for {
		listTeamResponse, err := gql.ListTeams(ctx, conn.client, pageSize, endCursor, true, &filters)
		if err != nil {
			plugin.Logger(ctx).Error("linear_team.listTeams", "api_error", err)
			return nil, err
		}

		for _, node := range listTeamResponse.Teams.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listTeamResponse.Teams.PageInfo.HasNextPage {
			break
		}
		endCursor = *listTeamResponse.Teams.PageInfo.EndCursor
	}

	return nil, nil
}

// HYDRATE FUNCTION

func getTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_team.getTeam", "connection_error", err)
		return nil, err
	}

	getTeamResponse, err := gql.GetTeam(ctx, conn.client, &id)
	if err != nil {
		plugin.Logger(ctx).Error("linear_team.getTeam", "api_error", err)
		return nil, err
	}

	return getTeamResponse.Team, nil
}

// Set the requested filter
func setTeamFilters(d *plugin.QueryData, ctx context.Context) gql.TeamFilter {
	var filter gql.TeamFilter
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
	if d.EqualsQuals["name"] != nil {
		name := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("name")),
		}
		filter.Name = name
	}
	if d.EqualsQuals["key"] != nil {
		key := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("key")),
		}
		filter.Key = key
	}

	return filter
}
