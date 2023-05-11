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
					Name:    "creator_id",
					Require: plugin.Optional,
				},
			},
		},
		// Get: &plugin.GetConfig{
		// 	KeyColumns: plugin.SingleColumn("id"),
		// 	Hydrate:    getProject,
		// },
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
				Description: "[Internal] The estimated start date of the project.",
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
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt. This is the same as the creation time if the entity hasn't been updated after creation.",
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
				Name:        "lead",
				Type:        proto.ColumnType_JSON,
				Description: "The project lead.",
			},
			{
				Name:        "creator_id",
				Type:        proto.ColumnType_STRING,
				Description: "The user who created the project.",
				Transform:   transform.FromField("Creator.Id"),
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
	var pageSize int = 100
	if d.QueryContext.Limit != nil {
		if int(*d.QueryContext.Limit) < pageSize {
			pageSize = int(*d.QueryContext.Limit)
		}
	}

	var filter *gql.ProjectFilter
	if d.EqualsQualString("creator_id") != "" {
		id := &gql.IDComparator{
			Eq: types.String(d.EqualsQualString("creator_id")),
		}
		creator := &gql.UserFilter{
			Id: id,
		}
		filter.Creator = creator
	}

	for {
		listProjectResponse, err := gql.ListProject(ctx, conn, pageSize, endCursor, filter)
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

// func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	id := d.EqualsQualString("id")

// 	// check if id is empty
// 	if id == "" {
// 		return nil, nil
// 	}

// 	conn, err := connect(ctx, d)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("linear_project.getProject", "connection_error", err)
// 		return nil, err
// 	}

// 	getProjectResponse, err := gql.GetProject(ctx, conn, id)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("linear_project.listProjects", "api_error", err)
// 		return nil, err
// 	}

// 	return getProjectResponse.Project, nil
// }
