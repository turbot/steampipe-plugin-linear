package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

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
				Type:        proto.ColumnType_STRING,
				Description: "A unique identifier for the team.",
			},
			{
				Name:        "archived_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the team was archived (if it has been).",
			},
			{
				Name:        "auto_archive_period",
				Type:        proto.ColumnType_INT,
				Description: "The period (in months) after which automatically closed and completed issues are automatically archived.",
			},
			{
				Name:        "auto_close_period",
				Type:        proto.ColumnType_INT,
				Description: "The period (in months) after which issues are automatically closed. Null/undefined means this feature is disabled.",
			},
			{
				Name:        "auto_close_state_id",
				Type:        proto.ColumnType_STRING,
				Description: "The canceled workflow state which auto closed issues will be set to. Defaults to the first canceled state.",
			},
			{
				Name:        "color",
				Type:        proto.ColumnType_STRING,
				Description: "The team's color.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the team was created.",
			},
			{
				Name:        "cycle_calender_url",
				Type:        proto.ColumnType_STRING,
				Description: "Calendar feed URL (iCal) for cycles.",
			},
			{
				Name:        "cycle_cooldown_time",
				Type:        proto.ColumnType_INT,
				Description: "The cooldown time after each cycle in weeks.",
			},
			{
				Name:        "cycle_duration",
				Type:        proto.ColumnType_INT,
				Description: "The duration of a cycle in weeks.",
			},
			{
				Name:        "cycle_issue_auto_assign_completed",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether completed issues are automatically assigned to the current cycle.",
			},
			{
				Name:        "cycle_issue_auto_assign_started",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether started issues are automatically assigned to the current cycle.",
			},
			{
				Name:        "cycle_lock_to_active",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether only issues with cycles in Active Issues are allowed.",
			},
			{
				Name:        "cycle_start_day",
				Type:        proto.ColumnType_STRING,
				Description: "The day of the week that a new cycle starts.",
			},
			{
				Name:        "cycles_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the team uses cycles.",
			},
			{
				Name:        "default_issue_estimate",
				Type:        proto.ColumnType_INT,
				Description: "The default estimate for unestimated issues.",
			},
			{
				Name:        "default_template_for_members_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the default template to use for new issues created by members of the team.",
			},
			{
				Name:        "default_template_for_non_members_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the default template to use for new issues created by non-members of the team.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The team's description.",
			},
			{
				Name:        "group_issue_history",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to group recent issue history entries.",
			},
			{
				Name:        "icon",
				Type:        proto.ColumnType_STRING,
				Description: "The team's icon.",
			},
			{
				Name:        "invite_hash",
				Type:        proto.ColumnType_STRING,
				Description: "A unique hash for the team to be used in invite URLs.",
			},
			{
				Name:        "issue_estimation_allow_zero",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to allow zeros in issues estimates.",
			},
			{
				Name:        "issue_estimation_extended",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to add additional points to the estimate scale.",
			},
			{
				Name:        "issue_estimation_type",
				Type:        proto.ColumnType_STRING,
				Description: "The issue estimation type to use.",
			},
			{
				Name:        "issue_ordering_no_priority_first",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether issues without priority should be sorted first.",
			},
			{
				Name:        "issue_sort_order_default_to_bottom",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to move issues to bottom of the column when changing state.",
			},
			{
				Name:        "key",
				Type:        proto.ColumnType_STRING,
				Description: "The team's unique key. The key is used in URLs.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The team's name.",
			},
			{
				Name:        "private",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the team is private or not.",
			},
			{
				Name:        "require_priority_to_leave_triage",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether an issue needs to have a priority set before leaving triage.",
			},
			{
				Name:        "slack_issue_comments",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to send new issue comment notifications to Slack.",
			},
			{
				Name:        "slack_issue_statuses",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to send new issue status updates to Slack.",
			},
			{
				Name:        "slack_new_issue",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to send new issue notifications to Slack.",
			},
			{
				Name:        "timezone",
				Type:        proto.ColumnType_STRING,
				Description: "The timezone of the team. Defaults to 'America/Los_Angeles'",
			},
			{
				Name:        "triage_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether triage mode is enabled for the team or not.",
			},
			{
				Name:        "upcoming_cycle_count",
				Type:        proto.ColumnType_DOUBLE,
				Description: "How many upcoming cycles to create.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "active_cycle",
				Type:        proto.ColumnType_JSON,
				Description: "Team's currently active cycle.",
			},
			{
				Name:        "default_issue_state",
				Type:        proto.ColumnType_JSON,
				Description: "The default workflow state into which issues are set when they are opened by team members.",
			},
			{
				Name:        "default_template_for_members",
				Type:        proto.ColumnType_JSON,
				Description: "The default template to use for new issues created by members of the team.",
			},
			{
				Name:        "default_template_for_non_members",
				Type:        proto.ColumnType_JSON,
				Description: "The default template to use for new issues created by non-members of the team.",
			},
			{
				Name:        "draft_workflow_state",
				Type:        proto.ColumnType_JSON,
				Description: "The workflow state into which issues are moved when a PR has been opened as draft.",
			},
			{
				Name:        "integrations_settings",
				Type:        proto.ColumnType_JSON,
				Description: "Settings for all integrations associated with that team.",
			},
			{
				Name:        "marked_as_duplicate_workflow_state",
				Type:        proto.ColumnType_JSON,
				Description: "The workflow state into which issues are moved when they are marked as a duplicate of another issue. Defaults to the first canceled state.",
			},
			{
				Name:        "organization",
				Type:        proto.ColumnType_JSON,
				Description: "The organization that the team is associated with.",
			},
			{
				Name:        "review_workflow_state",
				Type:        proto.ColumnType_JSON,
				Description: "The workflow state into which issues are moved when a review has been requested for the PR.",
			},
			{
				Name:        "start_workflow_state",
				Type:        proto.ColumnType_JSON,
				Description: "The workflow state into which issues are moved when a PR has been opened.",
			},
			{
				Name:        "triage_issue_state",
				Type:        proto.ColumnType_JSON,
				Description: "The workflow state into which issues are set when they are opened by non-team members or integrations if triage is enabled.",
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

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeCycle, includeIssueState, includeTemplateForMembers, includeTemplateForNonMembers, includeWorkflowState, includeIntegrationsSettings, includeDuplicateWorkflowState, includeOrganization, includeReviewWorkflowState, includeStartWorkflowState, includeTriageWorkflowState := true, true, true, true, true, true, true, true, true, true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "active_cycle":
			includeCycle = false
		case "default_issue_state":
			includeIssueState = false
		case "default_template_for_members":
			includeTemplateForMembers = false
		case "default_template_for_non_members":
			includeTemplateForNonMembers = false
		case "draft_workflow_state":
			includeWorkflowState = false
		case "integrations_settings":
			includeIntegrationsSettings = false
		case "marked_as_duplicate_workflow_state":
			includeDuplicateWorkflowState = false
		case "organization":
			includeOrganization = false
		case "review_workflow_state":
			includeReviewWorkflowState = false
		case "start_workflow_state":
			includeStartWorkflowState = false
		case "triage_issue_state":
			includeTriageWorkflowState = false
		}

	}

	// set the requested filters
	filters := setTeamFilters(d, ctx)

	for {
		listTeamResponse, err := gql.ListTeams(ctx, conn.client, pageSize, endCursor, true, &filters, &includeCycle, &includeIssueState, &includeTemplateForMembers, &includeTemplateForNonMembers, &includeWorkflowState, &includeIntegrationsSettings, &includeDuplicateWorkflowState, &includeOrganization, &includeReviewWorkflowState, &includeStartWorkflowState, &includeTriageWorkflowState)
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

func getTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeCycle, includeIssueState, includeTemplateForMembers, includeTemplateForNonMembers, includeWorkflowState, includeIntegrationsSettings, includeDuplicateWorkflowState, includeOrganization, includeReviewWorkflowState, includeStartWorkflowState, includeTriageWorkflowState := true, true, true, true, true, true, true, true, true, true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "active_cycle":
			includeCycle = false
		case "default_issue_state":
			includeIssueState = false
		case "default_template_for_members":
			includeTemplateForMembers = false
		case "default_template_for_non_members":
			includeTemplateForNonMembers = false
		case "draft_workflow_state":
			includeWorkflowState = false
		case "integrations_settings":
			includeIntegrationsSettings = false
		case "marked_as_duplicate_workflow_state":
			includeDuplicateWorkflowState = false
		case "organization":
			includeOrganization = false
		case "review_workflow_state":
			includeReviewWorkflowState = false
		case "start_workflow_state":
			includeStartWorkflowState = false
		case "triage_issue_state":
			includeTriageWorkflowState = false
		}

	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_team.getTeam", "connection_error", err)
		return nil, err
	}

	getTeamResponse, err := gql.GetTeam(ctx, conn.client, &id, &includeCycle, &includeIssueState, &includeTemplateForMembers, &includeTemplateForNonMembers, &includeWorkflowState, &includeIntegrationsSettings, &includeDuplicateWorkflowState, &includeOrganization, &includeReviewWorkflowState, &includeStartWorkflowState, &includeTriageWorkflowState)
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
