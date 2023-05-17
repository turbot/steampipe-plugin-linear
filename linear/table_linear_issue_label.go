package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinearIssueLabel(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_issue_label",
		Description: "Linear Issue Label",
		List: &plugin.ListConfig{
			Hydrate: listIssueLabels,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getIssueLabel,
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
				Description: "The label's color as a HEX string.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was created.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The label's description.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The label's name.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "team",
				Type:        proto.ColumnType_JSON,
				Description: "The team that the label is associated with. If null, the label is associated with the global workspace.",
			},
			{
				Name:        "issue_ids",
				Type:        proto.ColumnType_JSON,
				Description: "The issue ids associated with the label.",
				Transform:   transform.FromField("Issues.Nodes"),
			},
			{
				Name:        "creator",
				Type:        proto.ColumnType_JSON,
				Description: "The user who created the label.",
			},
			{
				Name:        "organization",
				Type:        proto.ColumnType_JSON,
				Description: "The organization associated with the label.",
			},
			{
				Name:        "parent",
				Type:        proto.ColumnType_JSON,
				Description: "The parent label.",
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

func listIssueLabels(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_issue_label.listIssueLabels", "connection_error", err)
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

	// set default pageSize for nested field issue ids
	var issuePageSize int = 50

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeCreator, includeOrganization, includeParent, includeTeam := true, true, true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "team":
			includeTeam = false
		case "organization":
			includeOrganization = false
		case "parent":
			includeParent = false
		case "creator":
			includeCreator = false
		}
	}

	// set the requested filters
	filters := setIssueLabelFilters(d, ctx)

	for {
		listIssueLabelResponse, err := gql.ListIssueLabels(ctx, conn.client, pageSize, issuePageSize, endCursor, true, &filters, &includeCreator, &includeOrganization, &includeParent, &includeTeam)
		if err != nil {
			plugin.Logger(ctx).Error("linear_issue_label.listIssueLabels", "api_error", err)
			return nil, err
		}
		for _, node := range listIssueLabelResponse.IssueLabels.Nodes {
			if *node.Issues.PageInfo.HasNextPage {
				endIssueCursor := *node.Issues.PageInfo.EndCursor
				for {
					getIssueIdsResponse, err := gql.GetIssueIds(ctx, conn.client, node.Id, issuePageSize, endIssueCursor, true)
					if err != nil {
						plugin.Logger(ctx).Error("linear_issue_label.listIssueLabels.GetIssueIds", "api_error", err)
						return nil, err
					}
					issueNodes := fetchIssueNodesFromList(getIssueIdsResponse.IssueLabel.Issues.Nodes)
					node.Issues.Nodes = append(node.Issues.Nodes, issueNodes...)
					if !*getIssueIdsResponse.IssueLabel.Issues.PageInfo.HasNextPage {
						break
					}
					endIssueCursor = *getIssueIdsResponse.IssueLabel.Issues.PageInfo.EndCursor
				}
			}
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listIssueLabelResponse.IssueLabels.PageInfo.HasNextPage {
			break
		}
		endCursor = *listIssueLabelResponse.IssueLabels.PageInfo.EndCursor
	}

	return nil, nil
}

func fetchIssueNodesFromList(nodes []*gql.GetIssuesNode) []*gql.ListIssuesNodes {
	var issueNodes []*gql.ListIssuesNodes
	for _, issueNode := range nodes {
		node := &gql.ListIssuesNodes{
			Id: issueNode.Id,
		}
		issueNodes = append(issueNodes, node)
	}
	return issueNodes
}

func getIssueLabel(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	var endCursor string

	// set default pageSize for nested field issue ids
	var issuePageSize int = 50

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeCreator, includeOrganization, includeParent, includeTeam := true, true, true, true
	for _, column := range d.QueryContext.Columns {
		switch column {
		case "team":
			includeTeam = false
		case "organization":
			includeOrganization = false
		case "parent":
			includeParent = false
		case "creator":
			includeCreator = false
		}
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_issue_label.getIssueLabel", "connection_error", err)
		return nil, err
	}

	getIssueLabelResponse, err := gql.GetIssueLabel(ctx, conn.client, &id, issuePageSize, true, &includeCreator, &includeOrganization, &includeParent, &includeTeam)
	if err != nil {
		plugin.Logger(ctx).Error("linear_issue_label.getIssueLabel", "api_error", err)
		return nil, err
	}
	if *getIssueLabelResponse.IssueLabel.Issues.PageInfo.HasNextPage {
		endCursor = *getIssueLabelResponse.IssueLabel.Issues.PageInfo.EndCursor
		for {
			getIssueIdsResponse, err := gql.GetIssueIds(ctx, conn.client, &id, issuePageSize, endCursor, true)
			if err != nil {
				plugin.Logger(ctx).Error("linear_issue_label.getIssueLabel.GetIssueIds", "api_error", err)
				return nil, err
			}
			issueNodes := fetchIssueNodesFromGet(getIssueIdsResponse.IssueLabel.Issues.Nodes)
			getIssueLabelResponse.IssueLabel.Issues.Nodes = append(getIssueLabelResponse.IssueLabel.Issues.Nodes, issueNodes...)
			if !*getIssueIdsResponse.IssueLabel.Issues.PageInfo.HasNextPage {
				break
			}
			endCursor = *getIssueIdsResponse.IssueLabel.Issues.PageInfo.EndCursor
		}
	}

	return getIssueLabelResponse.IssueLabel, nil
}

func fetchIssueNodesFromGet(nodes []*gql.GetIssuesNode) []*gql.GetIssueNode {
	var issueNodes []*gql.GetIssueNode
	for _, issueNode := range nodes {
		node := &gql.GetIssueNode{
			Id: issueNode.Id,
		}
		issueNodes = append(issueNodes, node)
	}
	return issueNodes
}

// Set the requested filter
func setIssueLabelFilters(d *plugin.QueryData, ctx context.Context) gql.IssueLabelFilter {
	var filter gql.IssueLabelFilter
	if d.EqualsQuals["name"] != nil {
		name := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("name")),
		}
		filter.Name = name
	}

	return filter
}
