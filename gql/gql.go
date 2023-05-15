package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

func ListIssue(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *IssueFilter, includeTeam *bool, includeCycle *bool, includeProject *bool, includeCreator *bool, includeAssignee *bool, includeSnoozedBy *bool, includeState *bool, includeParent *bool, includeProjectMilestone *bool) (*listIssueResponse, error) {
	return listIssue(ctx, client, first, after, includeArchived, filter, includeTeam, includeCycle, includeProject, includeCreator, includeAssignee, includeSnoozedBy, includeState, includeParent, includeProjectMilestone)
}

func ListAttachment(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *AttachmentFilter, includeCreator *bool, includeIssue *bool) (*listAttachmentResponse, error) {
	return listAttachment(ctx, client, first, after, includeArchived, filter, includeCreator, includeIssue)
}

func ListComment(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *CommentFilter, includeIssue *bool, includeParent *bool, includeUser *bool) (*listCommentResponse, error) {
	return listComment(ctx, client, first, after, includeArchived, filter, includeIssue, includeParent, includeUser)
}

func ListIntegration(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, includeCreator *bool, includeOrganization *bool, includeTeam *bool) (*listIntegrationResponse, error) {
	return listIntegration(ctx, client, first, after, includeArchived, includeCreator, includeOrganization, includeTeam)
}

func ListTeamMembership(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, includeTeam *bool, includeUser *bool) (*listTeamMembershipResponse, error) {
	return listTeamMembership(ctx, client, first, after, includeArchived, includeTeam, includeUser)
}

func ListProject(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *ProjectFilter, includeConvertedFromIssue *bool, includeIntegrationsSettings *bool, includeLead *bool, includeCreator *bool) (*listProjectResponse, error) {
	return listProject(ctx, client, first, after, includeArchived, filter, includeConvertedFromIssue, includeIntegrationsSettings, includeLead, includeCreator)
}

func ListUser(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *UserFilter, includeOrganization *bool) (*listUserResponse, error) {
	return listUser(ctx, client, first, after, includeArchived, filter, includeOrganization)
}

func ListTeam(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *TeamFilter, includeCycle *bool, includeIssueState *bool, includeTemplateForMembers *bool, includeTemplateForNonMembers *bool, includeWorkflowState *bool, includeIntegrationsSettings *bool, includeDuplicateWorkflowState *bool, includeOrganization *bool, includeReviewWorkflowState *bool, includeStartWorkflowState *bool, includeTriageWorkflowState *bool) (*listTeamResponse, error) {
	return listTeam(ctx, client, first, after, includeArchived, filter, includeCycle, includeIssueState, includeTemplateForMembers, includeTemplateForNonMembers, includeWorkflowState, includeIntegrationsSettings, includeDuplicateWorkflowState, includeOrganization, includeReviewWorkflowState, includeStartWorkflowState, includeTriageWorkflowState)
}

func ListWorkflowState(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *WorkflowStateFilter, includeTeam *bool) (*listWorkflowStateResponse, error) {
	return listWorkflowState(ctx, client, first, after, includeArchived, filter, includeTeam)
}

func ListIssueLabel(ctx context.Context, client graphql.Client, first int, after string, includeArchived bool, filter *IssueLabelFilter, includeCreator *bool, includeOrganization *bool, includeParent *bool, includeTeam *bool) (*listIssueLabelResponse, error) {
	return listIssueLabel(ctx, client, first, after, includeArchived, filter, includeCreator, includeOrganization, includeParent, includeTeam)
}

func ListOrganization(ctx context.Context, client graphql.Client, includeSubscription *bool) (*listOrganizationResponse, error) {
	return listOrganization(ctx, client, includeSubscription)
}
