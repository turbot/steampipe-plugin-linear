package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

func ListIssue(ctx context.Context, client graphql.Client, first int32, after string) (*listIssueResponse, error) {
	return listIssue(ctx, client, first, after)
}

func GetIssue(ctx context.Context, client graphql.Client, issueId string) (*getIssueResponse, error) {
	return getIssue(ctx, client, issueId)
}

func ListProject(ctx context.Context, client graphql.Client, first int32, after string) (*listProjectResponse, error) {
	return listProject(ctx, client, first, after)
}

func GetProject(ctx context.Context, client graphql.Client, issueId string) (*getProjectResponse, error) {
	return getProject(ctx, client, issueId)
}
