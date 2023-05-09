package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableLinearIssue(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_issue",
		Description: "Linear Issue",
		List: &plugin.ListConfig{
			Hydrate: listIssues,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getIssue,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the entity.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was created.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was updated. This is the same as the creation time if the entity hasn't been update after creation.",
			},
			{
				Name:        "archived_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was archived. Null if the entity has not been archived.",
			},
			{
				Name:        "number",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The issue's unique number.",
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "The issue's title.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The issue's description in markdown format.",
			},
			{
				Name:        "priority",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The priority of the issue.",
			},
			{
				Name:        "estimate",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The estimate of the complexity of the issue.",
			},
			{
				Name:        "sort_order",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The order of the item in relation to other items in the organization.",
			},
			{
				Name:        "started_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the issue was moved into started state.",
			},
			{
				Name:        "completed_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the issue was moved into completed state.",
			},
			{
				Name:        "canceled_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the issue was moved into canceled state.",
			},
			{
				Name:        "auto_closed_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the issue was automatically closed by the auto pruning process.",
			},
			{
				Name:        "auto_archived_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the issue was automatically archived by the auto pruning process.",
			},
			{
				Name:        "due_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date at which the issue is due.",
			},
			{
				Name:        "trashed",
				Type:        proto.ColumnType_BOOL,
				Description: "A flag that indicates whether the issue is in the trash bin.",
			},
			{
				Name:        "snoozed_util_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time until an issue will be snoozed in Triage view.",
			},
			{
				Name:        "team_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the team that the issue is associated with.",
			},
			{
				Name:        "cycle_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the cycle that the issue is associated with.",
			},
			{
				Name:        "project_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the project that the issue is associated with.",
			},
			{
				Name:        "previous_identifiers",
				Type:        proto.ColumnType_JSON,
				Description: "Previous identifiers of the issue if it has been moved between teams.",
			},
			{
				Name:        "creator_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the user who created the issue.",
			},
			{
				Name:        "assignee_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the user to whom the issue is assigned to.",
			},
			{
				Name:        "snoozed_by_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the user who snoozed the issue.",
			},
			{
				Name:        "state_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the workflow state that the issue is associated with.",
			},
			{
				Name:        "parent_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the parent of the issue.",
			},
			{
				Name:        "sub_issue_sort_order",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The unique identifier of the parent of the issue.",
			},
			{
				Name:        "priority_label",
				Type:        proto.ColumnType_STRING,
				Description: "Label for the priority.",
			},
			{
				Name:        "identifier",
				Type:        proto.ColumnType_STRING,
				Description: "Issue's human readable identifier (e.g. ENG-123).",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "Issue URL.",
			},
			{
				Name:        "branch_name",
				Type:        proto.ColumnType_STRING,
				Description: "Suggested branch name for the issue.",
			},
			{
				Name:        "customer_ticket_count",
				Type:        proto.ColumnType_INT,
				Description: "Returns the number of Attachment resources which are created by customer support ticketing systems (e.g. Zendesk).",
			},
			{
				Name:        "subscriber_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the users who are subscribed to the issue.",
			},
			{
				Name:        "children_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the children of the issue.",
			},
			{
				Name:        "comment_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the comments associated with the issue.",
			},
			{
				Name:        "history_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the history entries associated with the issue.",
			},
			{
				Name:        "label_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the labels associated with the issue.",
			},
			{
				Name:        "integration_resource_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the integration resources for this issue.",
			},
			{
				Name:        "relation_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the relations associated with this issue.",
			},
			{
				Name:        "inverse_relation_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the inverse relations associated with this issue.",
			},
			{
				Name:        "attachment_ids",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unique identifiers of the attachments associated with the issue.",
			},
		},
	}
}

func listIssues(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_issue.listIssues", "connection_error", err)
		return nil, err
	}
	var endCursor string
	var pageSize int32 = 100
	if d.QueryContext.Limit != nil {
		if int32(*d.QueryContext.Limit) < pageSize {
			pageSize = int32(*d.QueryContext.Limit)
		}
	}

	for {
		listIssueResponse, err := gql.ListIssue(ctx, conn, pageSize, endCursor)
		if err != nil {
			plugin.Logger(ctx).Error("linear_issue.listIssues", "api_error", err)
			return nil, err
		}
		for _, node := range listIssueResponse.Issues.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !listIssueResponse.Issues.PageInfo.HasNextPage {
			break
		}
		endCursor = listIssueResponse.Issues.PageInfo.EndCursor
	}

	return nil, nil
}

func getIssue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_issue.getIssue", "connection_error", err)
		return nil, err
	}

	getIssueResponse, err := gql.GetIssue(ctx, conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("linear_issue.listIssues", "api_error", err)
		return nil, err
	}

	return getIssueResponse.Issue, nil
}
