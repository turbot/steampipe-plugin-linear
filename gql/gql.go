package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type ListIssuesNodes = listIssueLabelsIssueLabelsIssueLabelConnectionNodesIssueLabelIssuesIssueConnectionNodesIssue
type GetIssuesNode = getIssueIdsIssueLabelIssuesIssueConnectionNodesIssue
type GetIssueNode = getIssueLabelIssueLabelIssuesIssueConnectionNodesIssue

func ListIssues(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *IssueFilter) (*listIssuesResponse, error) {
	return listIssues(ctx, client, first, after, includeArchived, filter)
}

func GetIssue(ctx context.Context, client graphql.Client, id *string) (*getIssueResponse, error) {
	return getIssue(ctx, client, id)
}

func ListAttachments(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *AttachmentFilter) (*listAttachmentsResponse, error) {
	return listAttachments(ctx, client, first, after, includeArchived, filter)
}

func GetAttachment(ctx context.Context, client graphql.Client, id *string) (*getAttachmentResponse, error) {
	return getAttachment(ctx, client, id)
}

func ListComments(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *CommentFilter) (*listCommentsResponse, error) {
	return listComments(ctx, client, first, after, includeArchived, filter)
}

func GetComment(ctx context.Context, client graphql.Client, id *string) (*getCommentResponse, error) {
	return getComment(ctx, client, id)
}

func ListIntegrations(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool) (*listIntegrationsResponse, error) {
	return listIntegrations(ctx, client, first, after, includeArchived)
}

func GetIntegration(ctx context.Context, client graphql.Client, id *string) (*getIntegrationResponse, error) {
	return getIntegration(ctx, client, id)
}

func ListTeamMemberships(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool) (*listTeamMembershipsResponse, error) {
	return listTeamMemberships(ctx, client, first, after, includeArchived)
}

func GetTeamMembership(ctx context.Context, client graphql.Client, id *string) (*getTeamMembershipResponse, error) {
	return getTeamMembership(ctx, client, id)
}

func ListProjects(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *ProjectFilter) (*listProjectsResponse, error) {
	return listProjects(ctx, client, first, after, includeArchived, filter)
}

func GetProject(ctx context.Context, client graphql.Client, id *string) (*getProjectResponse, error) {
	return getProject(ctx, client, id)
}

func ListUsers(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *UserFilter) (*listUsersResponse, error) {
	return listUsers(ctx, client, first, after, includeArchived, filter)
}

func GetUser(ctx context.Context, client graphql.Client, id *string) (*getUserResponse, error) {
	return getUser(ctx, client, id)
}

func ListTeams(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *TeamFilter) (*listTeamsResponse, error) {
	return listTeams(ctx, client, first, after, includeArchived, filter)
}

func GetTeam(ctx context.Context, client graphql.Client, id *string) (*getTeamResponse, error) {
	return getTeam(ctx, client, id)
}

func ListIssueLabels(ctx context.Context, client graphql.Client, first int, firstIssue int, after string, includeArchived bool, filter *IssueLabelFilter) (*listIssueLabelsResponse, error) {
	return listIssueLabels(ctx, client, first, firstIssue, after, includeArchived, filter)
}

func GetIssueLabel(ctx context.Context, client graphql.Client, id *string, first int, includeArchived bool) (*getIssueLabelResponse, error) {
	return getIssueLabel(ctx, client, id, first, includeArchived)
}

func GetIssueIds(ctx context.Context, client graphql.Client, id *string, first int, after string, includeArchived bool) (*getIssueIdsResponse, error) {
	return getIssueIds(ctx, client, id, first, after, includeArchived)
}

func GetOrganization(ctx context.Context, client graphql.Client) (*getOrganizationResponse, error) {
	return getOrganization(ctx, client)
}
