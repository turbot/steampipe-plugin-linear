package linear

import (
	"context"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-linear/gql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableLinearIssue(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_issue",
		Description: "Linear Issue",
		List: &plugin.ListConfig{
			Hydrate: listIssues,
			KeyColumns: []*plugin.KeyColumn{
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
					Name:      "number",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:    "title",
					Require: plugin.Optional,
				},
				{
					Name:      "priority",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "started_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "completed_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "canceled_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "auto_closed_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "auto_archived_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "due_date",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "snoozed_until_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getIssue,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier of the entity.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time at which the entity was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_at",
				Description: "The last time at which the entity was meaningfully updated. This is the same as the creation time if the entity hasn't been updated after creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "archived_at",
				Description: "The time at which the entity was archived. Null if the entity has not been archived.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "number",
				Description: "The issue's unique number.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "description",
				Description: "The issue's description in markdown format.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "priority",
				Description: "The priority of the issue. 0 = No priority, 1 = Urgent, 2 = High, 3 = Normal, 4 = Low.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "estimate",
				Description: "The estimate of the complexity of the issue.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "sort_order",
				Description: "The order of the item in relation to other items in the organization.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "started_at",
				Description: "The time at which the issue was moved into started state.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "completed_at",
				Description: "The time at which the issue was moved into completed state.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "canceled_at",
				Description: "The time at which the issue was moved into canceled state.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "auto_closed_at",
				Description: "The time at which the issue was automatically closed by the auto pruning process.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "auto_archived_at",
				Description: "The time at which the issue was automatically archived by the auto pruning process.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "due_date",
				Description: "The date at which the issue is due.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "trashed",
				Description: "A flag that indicates whether the issue is in the trash bin.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "snoozed_until_at",
				Description: "The time until an issue will be snoozed in Triage view.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "previous_identifiers",
				Description: "Previous identifiers of the issue if it has been moved between teams.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "sub_issue_sort_order",
				Description: "The order of the item in the sub-issue list. Only set if the issue has a parent.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "priority_label",
				Description: "Label for the priority.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "identifier",
				Description: "Issue's human readable identifier (e.g. ENG-123).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "url",
				Description: "Issue URL.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "branch_name",
				Description: "Suggested branch name for the issue.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer_ticket_count",
				Description: "Returns the number of Attachment resources which are created by customer support ticketing systems (e.g. Zendesk).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "team",
				Description: "The team that the issue is associated with.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "cycle",
				Description: "The cycle that the issue is associated with.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "project",
				Description: "The project that the issue is associated with.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "creator",
				Description: "The user who created the issue.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "assignee",
				Description: "The user to whom the issue is assigned to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "snoozed_by",
				Description: "The user who snoozed the issue.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "state",
				Description: "The workflow state that the issue is associated with.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "parent",
				Description: "The parent of the issue.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "project_milestone",
				Description: "The projectMilestone that the issue is associated with.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The issue's title.",
				Type:        proto.ColumnType_STRING,
			},
		}),
	}
}

// LIST FUNCTION

func listIssues(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_issue.listIssues", "connection_error", err)
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
	filters := setIssueFilters(d, ctx)

	for {
		listIssueResponse, err := gql.ListIssues(ctx, conn.client, pageSize, endCursor, true, &filters)
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
		if !*listIssueResponse.Issues.PageInfo.HasNextPage {
			break
		}
		endCursor = *listIssueResponse.Issues.PageInfo.EndCursor
	}

	return nil, nil
}

// HYDRATE FUNCTION

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

	getIssueResponse, err := gql.GetIssue(ctx, conn.client, &id)
	if err != nil {
		plugin.Logger(ctx).Error("linear_issue.listIssues", "api_error", err)
		return nil, err
	}

	return getIssueResponse.Issue, nil
}

// Set the requested filter
func setIssueFilters(d *plugin.QueryData, ctx context.Context) gql.IssueFilter {
	var filter gql.IssueFilter
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
	if d.Quals["number"] != nil {
		numberCom := &gql.NumberComparator{}
		for _, q := range d.Quals["number"].Quals {
			number := types.Float64(q.Value.GetDoubleValue())
			switch q.Operator {
			case "=":
				numberCom.Eq = number
			case ">":
				numberCom.Gt = number
			case ">=":
				numberCom.Gte = number
			case "<":
				numberCom.Lt = number
			case "<=":
				numberCom.Lte = number
			}
		}

		filter.Number = numberCom
	}
	if d.EqualsQuals["title"] != nil {
		title := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("title")),
		}
		filter.Title = title
	}
	if d.Quals["priority"] != nil {
		priorityCom := &gql.NullableNumberComparator{}
		for _, q := range d.Quals["priority"].Quals {
			priority := types.Float64(q.Value.GetDoubleValue())
			switch q.Operator {
			case "=":
				priorityCom.Eq = priority
			case ">":
				priorityCom.Gt = priority
			case ">=":
				priorityCom.Gte = priority
			case "<":
				priorityCom.Lt = priority
			case "<=":
				priorityCom.Lte = priority
			}
		}

		filter.Priority = priorityCom
	}
	if d.Quals["started_at"] != nil {
		startedAt := &gql.NullableDateComparator{}
		for _, q := range d.Quals["started_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				startedAt.Eq = timestamp
			case ">":
				startedAt.Gt = timestamp
			case ">=":
				startedAt.Gte = timestamp
			case "<":
				startedAt.Lt = timestamp
			case "<=":
				startedAt.Lte = timestamp
			}
		}
		filter.StartedAt = startedAt
	}
	if d.Quals["completed_at"] != nil {
		completedAt := &gql.NullableDateComparator{}
		for _, q := range d.Quals["completed_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				completedAt.Eq = timestamp
			case ">":
				completedAt.Gt = timestamp
			case ">=":
				completedAt.Gte = timestamp
			case "<":
				completedAt.Lt = timestamp
			case "<=":
				completedAt.Lte = timestamp
			}
		}
		filter.CompletedAt = completedAt
	}
	if d.Quals["canceled_at"] != nil {
		canceledAt := &gql.NullableDateComparator{}
		for _, q := range d.Quals["canceled_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				canceledAt.Eq = timestamp
			case ">":
				canceledAt.Gt = timestamp
			case ">=":
				canceledAt.Gte = timestamp
			case "<":
				canceledAt.Lt = timestamp
			case "<=":
				canceledAt.Lte = timestamp
			}
		}
		filter.CanceledAt = canceledAt
	}
	if d.Quals["auto_closed_at"] != nil {
		autoClosedAt := &gql.NullableDateComparator{}
		for _, q := range d.Quals["auto_closed_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				autoClosedAt.Eq = timestamp
			case ">":
				autoClosedAt.Gt = timestamp
			case ">=":
				autoClosedAt.Gte = timestamp
			case "<":
				autoClosedAt.Lt = timestamp
			case "<=":
				autoClosedAt.Lte = timestamp
			}
		}
		filter.AutoClosedAt = autoClosedAt
	}
	if d.Quals["auto_archived_at"] != nil {
		autoArchivedAt := &gql.NullableDateComparator{}
		for _, q := range d.Quals["auto_archived_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				autoArchivedAt.Eq = timestamp
			case ">":
				autoArchivedAt.Gt = timestamp
			case ">=":
				autoArchivedAt.Gte = timestamp
			case "<":
				autoArchivedAt.Lt = timestamp
			case "<=":
				autoArchivedAt.Lte = timestamp
			}
		}
		filter.AutoArchivedAt = autoArchivedAt
	}
	if d.Quals["due_date"] != nil {
		dueDate := &gql.NullableTimelessDateComparator{}
		for _, q := range d.Quals["due_date"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				dueDate.Eq = timestamp
			case ">":
				dueDate.Gt = timestamp
			case ">=":
				dueDate.Gte = timestamp
			case "<":
				dueDate.Lt = timestamp
			case "<=":
				dueDate.Lte = timestamp
			}
		}
		filter.DueDate = dueDate
	}
	if d.Quals["snoozed_until_at"] != nil {
		snoozedUntilAt := &gql.NullableDateComparator{}
		for _, q := range d.Quals["snoozed_until_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				snoozedUntilAt.Eq = timestamp
			case ">":
				snoozedUntilAt.Gt = timestamp
			case ">=":
				snoozedUntilAt.Gte = timestamp
			case "<":
				snoozedUntilAt.Lt = timestamp
			case "<=":
				snoozedUntilAt.Lte = timestamp
			}
		}
		filter.SnoozedUntilAt = snoozedUntilAt
	}
	return filter
}
