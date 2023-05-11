package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

func ListIssue(ctx context.Context, client graphql.Client, first int, after string, filter *IssueFilter, includeTeam *bool, includeCycle *bool, includeProject *bool, includeCreator *bool, includeAssignee *bool, includeSnoozedBy *bool, includeState *bool, includeParent *bool, includeProjectMilestone *bool) (*listIssueResponse, error) {
	return listIssue(ctx, client, first, after, filter, includeTeam, includeCycle, includeProject, includeCreator, includeAssignee, includeSnoozedBy, includeState, includeParent, includeProjectMilestone)
}

func ListProject(ctx context.Context, client graphql.Client, first int, after string, filter *ProjectFilter) (*listProjectResponse, error) {
	return listProject(ctx, client, first, after, filter)
}
