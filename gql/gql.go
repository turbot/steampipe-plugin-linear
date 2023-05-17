package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type ListIssuesNodes = listIssueLabelsIssueLabelsIssueLabelConnectionNodesIssueLabelIssuesIssueConnectionNodesIssue
type GetIssuesNode = getIssueIdsIssueLabelIssuesIssueConnectionNodesIssue
type GetIssueNode = getIssueLabelIssueLabelIssuesIssueConnectionNodesIssue

func ListIssues(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *IssueFilter, includeTeam *bool, includeCycle *bool, includeProject *bool, includeCreator *bool, includeAssignee *bool, includeSnoozedBy *bool, includeState *bool, includeParent *bool, includeProjectMilestone *bool) (*listIssuesResponse, error) {
	return listIssues(ctx, client, first, after, includeArchived, filter, includeTeam, includeCycle, includeProject, includeCreator, includeAssignee, includeSnoozedBy, includeState, includeParent, includeProjectMilestone)
}

func GetIssue(ctx context.Context, client graphql.Client, id *string, includeTeam *bool, includeCycle *bool, includeProject *bool, includeCreator *bool, includeAssignee *bool, includeSnoozedBy *bool, includeState *bool, includeParent *bool, includeProjectMilestone *bool) (*getIssueResponse, error) {
	return getIssue(ctx, client, id, includeTeam, includeCycle, includeProject, includeCreator, includeAssignee, includeSnoozedBy, includeState, includeParent, includeProjectMilestone)
}

func ListAttachments(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *AttachmentFilter, includeCreator *bool, includeIssue *bool) (*listAttachmentsResponse, error) {
	return listAttachments(ctx, client, first, after, includeArchived, filter, includeCreator, includeIssue)
}

func GetAttachment(ctx context.Context, client graphql.Client, id *string, includeCreator *bool, includeIssue *bool) (*getAttachmentResponse, error) {
	return getAttachment(ctx, client, id, includeCreator, includeIssue)
}

func ListComments(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *CommentFilter, includeIssue *bool, includeParent *bool, includeUser *bool) (*listCommentsResponse, error) {
	return listComments(ctx, client, first, after, includeArchived, filter, includeIssue, includeParent, includeUser)
}

func GetComment(ctx context.Context, client graphql.Client, id *string, includeIssue *bool, includeParent *bool, includeUser *bool) (*getCommentResponse, error) {
	return getComment(ctx, client, id, includeIssue, includeParent, includeUser)
}

func ListIntegrations(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, includeCreator *bool, includeOrganization *bool, includeTeam *bool) (*listIntegrationsResponse, error) {
	return listIntegrations(ctx, client, first, after, includeArchived, includeCreator, includeOrganization, includeTeam)
}

func GetIntegration(ctx context.Context, client graphql.Client, id *string, includeCreator *bool, includeOrganization *bool, includeTeam *bool) (*getIntegrationResponse, error) {
	return getIntegration(ctx, client, id, includeCreator, includeOrganization, includeTeam)
}

func ListTeamMemberships(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, includeTeam *bool, includeUser *bool) (*listTeamMembershipsResponse, error) {
	return listTeamMemberships(ctx, client, first, after, includeArchived, includeTeam, includeUser)
}

func GetTeamMembership(ctx context.Context, client graphql.Client, id *string, includeTeam *bool, includeUser *bool) (*getTeamMembershipResponse, error) {
	return getTeamMembership(ctx, client, id, includeTeam, includeUser)
}

func ListProjects(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *ProjectFilter, includeConvertedFromIssue *bool, includeIntegrationsSettings *bool, includeLead *bool, includeCreator *bool) (*listProjectsResponse, error) {
	return listProjects(ctx, client, first, after, includeArchived, filter, includeConvertedFromIssue, includeIntegrationsSettings, includeLead, includeCreator)
}

func GetProject(ctx context.Context, client graphql.Client, id *string, includeConvertedFromIssue *bool, includeIntegrationsSettings *bool, includeLead *bool, includeCreator *bool) (*getProjectResponse, error) {
	return getProject(ctx, client, id, includeConvertedFromIssue, includeIntegrationsSettings, includeLead, includeCreator)
}

func ListUsers(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *UserFilter, includeOrganization *bool) (*listUsersResponse, error) {
	return listUsers(ctx, client, first, after, includeArchived, filter, includeOrganization)
}

func GetUser(ctx context.Context, client graphql.Client, id *string, includeOrganization *bool) (*getUserResponse, error) {
	return getUser(ctx, client, id, includeOrganization)
}

func ListTeams(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *TeamFilter, includeCycle *bool, includeIssueState *bool, includeTemplateForMembers *bool, includeTemplateForNonMembers *bool, includeWorkflowState *bool, includeIntegrationsSettings *bool, includeDuplicateWorkflowState *bool, includeOrganization *bool, includeReviewWorkflowState *bool, includeStartWorkflowState *bool, includeTriageWorkflowState *bool) (*listTeamsResponse, error) {
	return listTeams(ctx, client, first, after, includeArchived, filter, includeCycle, includeIssueState, includeTemplateForMembers, includeTemplateForNonMembers, includeWorkflowState, includeIntegrationsSettings, includeDuplicateWorkflowState, includeOrganization, includeReviewWorkflowState, includeStartWorkflowState, includeTriageWorkflowState)
}

func GetTeam(ctx context.Context, client graphql.Client, id *string, includeCycle *bool, includeIssueState *bool, includeTemplateForMembers *bool, includeTemplateForNonMembers *bool, includeWorkflowState *bool, includeIntegrationsSettings *bool, includeDuplicateWorkflowState *bool, includeOrganization *bool, includeReviewWorkflowState *bool, includeStartWorkflowState *bool, includeTriageWorkflowState *bool) (*getTeamResponse, error) {
	return getTeam(ctx, client, id, includeCycle, includeIssueState, includeTemplateForMembers, includeTemplateForNonMembers, includeWorkflowState, includeIntegrationsSettings, includeDuplicateWorkflowState, includeOrganization, includeReviewWorkflowState, includeStartWorkflowState, includeTriageWorkflowState)
}

func ListIssueLabels(ctx context.Context, client graphql.Client, first int, firstIssue int, after string, includeArchived bool, filter *IssueLabelFilter, includeCreator *bool, includeOrganization *bool, includeParent *bool, includeTeam *bool) (*listIssueLabelsResponse, error) {
	return listIssueLabels(ctx, client, first, firstIssue, after, includeArchived, filter, includeCreator, includeOrganization, includeParent, includeTeam)
}

func GetIssueLabel(ctx context.Context, client graphql.Client, id *string, first int, includeArchived bool, includeCreator *bool, includeOrganization *bool, includeParent *bool, includeTeam *bool) (*getIssueLabelResponse, error) {
	return getIssueLabel(ctx, client, id, first, includeArchived, includeCreator, includeOrganization, includeParent, includeTeam)
}

func GetIssueIds(ctx context.Context, client graphql.Client, id *string, first int, after string, includeArchived bool) (*getIssueIdsResponse, error) {
	return getIssueIds(ctx, client, id, first, after, includeArchived)
}

func GetOrganization(ctx context.Context, client graphql.Client, includeSubscription *bool) (*getOrganizationResponse, error) {
	return getOrganization(ctx, client, includeSubscription)
}
