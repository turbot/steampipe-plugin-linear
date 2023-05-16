package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinearWorkflowState(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_workflow_state",
		Description: "Linear Workflow State",
		List: &plugin.ListConfig{
			Hydrate: listWorkflowStates,
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
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "position",
					Require: plugin.Optional,
				},
				{
					Name:    "type",
					Require: plugin.Optional,
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
			Hydrate:    getWorkflowState,
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
				Name:        "color",
				Type:        proto.ColumnType_STRING,
				Description: "The state's UI color as a HEX string.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was created.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description of the state.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The state's name.",
			},
			{
				Name:        "position",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The position of the state in the team flow.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the state.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e., for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "team",
				Type:        proto.ColumnType_JSON,
				Description: "The team to which this state belongs.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The issue label's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listWorkflowStates(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_workflow_state.listWorkflowStates", "connection_error", err)
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
	includeTeam := true
	if helpers.StringSliceContains(d.QueryContext.Columns, "team") {
		includeTeam = false
	}

	// set the requested filters
	filters := setWorkflowStateFilters(d, ctx)

	for {
		listWorkflowStateResponse, err := gql.ListWorkflowStates(ctx, conn, pageSize, endCursor, true, &filters, &includeTeam)
		if err != nil {
			plugin.Logger(ctx).Error("linear_workflow_state.listWorkflowStates", "api_error", err)
			return nil, err
		}

		for _, node := range listWorkflowStateResponse.WorkflowStates.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listWorkflowStateResponse.WorkflowStates.PageInfo.HasNextPage {
			break
		}
		endCursor = *listWorkflowStateResponse.WorkflowStates.PageInfo.EndCursor
	}

	return nil, nil
}

func getWorkflowState(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeTeam := true
	if helpers.StringSliceContains(d.QueryContext.Columns, "team") {
		includeTeam = false
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_workflow_state.getWorkflowState", "connection_error", err)
		return nil, err
	}

	getWorkflowStateResponse, err := gql.GetWorkflowState(ctx, conn, &id, &includeTeam)
	if err != nil {
		plugin.Logger(ctx).Error("linear_workflow_state.getWorkflowState", "api_error", err)
		return nil, err
	}

	return getWorkflowStateResponse.WorkflowState, nil
}

// Set the requested filter
func setWorkflowStateFilters(d *plugin.QueryData, ctx context.Context) gql.WorkflowStateFilter {
	var filter gql.WorkflowStateFilter
	if d.EqualsQuals["id"] != nil {
		id := &gql.IDComparator{
			Eq: types.String(d.EqualsQualString("id")),
		}
		filter.Id = id
	}
	if d.EqualsQuals["name"] != nil {
		name := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("name")),
		}
		filter.Name = name
	}
	if d.EqualsQuals["type"] != nil {
		typefilter := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("type")),
		}
		filter.Type = typefilter
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
	if d.Quals["position"] != nil {
		positionCon := &gql.NumberComparator{}
		for _, q := range d.Quals["position"].Quals {
			position := types.Float64(q.Value.GetDoubleValue())
			switch q.Operator {
			case "=":
				positionCon.Eq = position
			case ">":
				positionCon.Gt = position
			case ">=":
				positionCon.Gte = position
			case "<":
				positionCon.Lt = position
			case "<=":
				positionCon.Lte = position
			}
		}

		filter.Position = positionCon
	}

	return filter
}
