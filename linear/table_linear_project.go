package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinearProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_project",
		Description: "Linear Project",
		List: &plugin.ListConfig{
			Hydrate: listProjects,
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
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "state",
					Require: plugin.Optional,
				},
				{
					Name:    "slug_id",
					Require: plugin.Optional,
				},
				{
					Name:      "start_date",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "target_date",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProject,
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
				Name:        "auto_archived_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the project was automatically archived by the auto pruning process.",
			},
			{
				Name:        "canceled_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the project was moved into canceled state.",
			},
			{
				Name:        "color",
				Type:        proto.ColumnType_STRING,
				Description: "The project's color.",
			},
			{
				Name:        "completed_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the project was moved into completed state.",
			},
			{
				Name:        "completed_issue_count_history",
				Type:        proto.ColumnType_JSON,
				Description: "The number of completed issues in the project after each week.",
			},
			{
				Name:        "completed_scope_history",
				Type:        proto.ColumnType_JSON,
				Description: "The number of completed estimation points after each week.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was created.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The project's description.",
			},
			{
				Name:        "icon",
				Type:        proto.ColumnType_STRING,
				Description: "The icon of the project.",
			},
			{
				Name:        "in_progress_scope_history",
				Type:        proto.ColumnType_JSON,
				Description: "The number of in progress estimation points after each week.",
			},
			{
				Name:        "issue_count_history",
				Type:        proto.ColumnType_JSON,
				Description: "The total number of issues in the project after each week.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The project's name.",
			},
			{
				Name:        "progress",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The overall progress of the project. This is the (completed estimate points + 0.25 * in progress estimate points) / total estimate points.",
			},
			{
				Name:        "project_update_reminders_paused_until_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time until which project update reminders are paused.",
			},
			{
				Name:        "scope",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The overall scope (total estimate points) of the project.",
			},
			{
				Name:        "scope_history",
				Type:        proto.ColumnType_JSON,
				Description: "The total number of estimation points after each week.",
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
				Name:        "slug_id",
				Type:        proto.ColumnType_STRING,
				Description: "The project's unique URL slug.",
			},
			{
				Name:        "sort_order",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The sort order for the project within the organization.",
			},
			{
				Name:        "start_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The estimated start date of the project.",
			},
			{
				Name:        "started_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the project was moved into started state.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the state.",
			},
			{
				Name:        "target_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The estimated completion date of the project.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "Project URL.",
			},
			{
				Name:        "converted_from_issue",
				Type:        proto.ColumnType_JSON,
				Description: "The project was created based on this issue.",
			},
			{
				Name:        "integrations_settings",
				Type:        proto.ColumnType_JSON,
				Description: "Settings for all integrations associated with that project.",
			},
			{
				Name:        "creator",
				Type:        proto.ColumnType_JSON,
				Description: "The user who created the project.",
			},
			{
				Name:        "lead",
				Type:        proto.ColumnType_JSON,
				Description: "The project lead.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The project's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_project.listProjects", "connection_error", err)
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
	includeConvertedFromIssue, includeIntegrationsSettings, includeLead, includeCreator := true, true, true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "converted_from_issue":
			includeConvertedFromIssue = false
		case "integrations_settings":
			includeIntegrationsSettings = false
		case "lead":
			includeLead = false
		case "creator":
			includeCreator = false
		}
	}

	// set the requested filters
	filters := setProjectFilters(d, ctx)

	for {
		listProjectResponse, err := gql.ListProjects(ctx, conn, pageSize, endCursor, true, &filters, &includeConvertedFromIssue, &includeIntegrationsSettings, &includeLead, &includeCreator)
		if err != nil {
			plugin.Logger(ctx).Error("linear_project.listProjects", "api_error", err)
			return nil, err
		}
		for _, node := range listProjectResponse.Projects.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listProjectResponse.Projects.PageInfo.HasNextPage {
			break
		}
		endCursor = *listProjectResponse.Projects.PageInfo.EndCursor
	}

	return nil, nil
}

func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeConvertedFromIssue, includeIntegrationsSettings, includeLead, includeCreator := true, true, true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "converted_from_issue":
			includeConvertedFromIssue = false
		case "integrations_settings":
			includeIntegrationsSettings = false
		case "lead":
			includeLead = false
		case "creator":
			includeCreator = false
		}
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_project.getProject", "connection_error", err)
		return nil, err
	}

	getProjectResponse, err := gql.GetProject(ctx, conn, &id, &includeConvertedFromIssue, &includeIntegrationsSettings, &includeLead, &includeCreator)
	if err != nil {
		plugin.Logger(ctx).Error("linear_project.getProject", "api_error", err)
		return nil, err
	}

	return getProjectResponse.Project, nil
}

// Set the requested filter
func setProjectFilters(d *plugin.QueryData, ctx context.Context) gql.ProjectFilter {
	var filter gql.ProjectFilter
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
	if d.EqualsQuals["name"] != nil {
		name := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("name")),
		}
		filter.Name = name
	}
	if d.EqualsQuals["state"] != nil {
		state := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("state")),
		}
		filter.State = state
	}
	if d.EqualsQuals["slug_id"] != nil {
		slug_id := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("slug_id")),
		}
		filter.SlugId = slug_id
	}
	if d.Quals["start_date"] != nil {
		startDate := &gql.NullableDateComparator{}
		for _, q := range d.Quals["start_date"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				startDate.Eq = timestamp
			case ">":
				startDate.Gt = timestamp
			case ">=":
				startDate.Gte = timestamp
			case "<":
				startDate.Lt = timestamp
			case "<=":
				startDate.Lte = timestamp
			}
		}
		filter.StartDate = startDate
	}
	if d.Quals["target_date"] != nil {
		targetDate := &gql.NullableDateComparator{}
		for _, q := range d.Quals["target_date"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				targetDate.Eq = timestamp
			case ">":
				targetDate.Gt = timestamp
			case ">=":
				targetDate.Gte = timestamp
			case "<":
				targetDate.Lt = timestamp
			case "<=":
				targetDate.Lte = timestamp
			}
		}
		filter.TargetDate = targetDate
	}

	return filter
}
