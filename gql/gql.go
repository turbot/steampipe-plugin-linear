package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

func ListIssue(ctx context.Context, client graphql.Client, first int, after string, filter *IssueFilter, includeTeam *bool, includeCycle *bool, includeProject *bool, includeCreator *bool, includeAssignee *bool, includeSnoozedBy *bool, includeState *bool, includeParent *bool, includeProjectMilestone *bool) (*listIssueResponse, error) {
	return listIssue(ctx, client, first, after, filter, includeTeam, includeCycle, includeProject, includeCreator, includeAssignee, includeSnoozedBy, includeState, includeParent, includeProjectMilestone)
}

func ListAttachment(ctx context.Context, client graphql.Client, first int, after string, filter *AttachmentFilter, includeCreator *bool, includeIssue *bool) (*listAttachmentResponse, error) {
	return listAttachment(ctx, client, first, after, filter, includeCreator, includeIssue)
}

func ListComment(ctx context.Context, client graphql.Client, first int, after string, filter *CommentFilter, includeIssue *bool, includeParent *bool, includeUser *bool) (*listCommentResponse, error) {
	return listComment(ctx, client, first, after, filter, includeIssue, includeParent, includeUser)
}

func ListIntegration(ctx context.Context, client graphql.Client, first int, after string, includeCreator *bool, includeOrganization *bool, includeTeam *bool) (*listIntegrationResponse, error) {
	return listIntegration(ctx, client, first, after, includeCreator, includeOrganization, includeTeam)
}

func ListTeamMembership(ctx context.Context, client graphql.Client, first int, after string, includeTeam *bool, includeUser *bool) (*listTeamMembershipResponse, error) {
	return listTeamMembership(ctx, client, first, after, includeTeam, includeUser)
}

func ListProject(ctx context.Context, client graphql.Client, first int, after string, filter *ProjectFilter, includeConvertedFromIssue *bool, includeIntegrationsSettings *bool, includeLead *bool, includeCreator *bool) (*listProjectResponse, error) {
	return listProject(ctx, client, first, after, filter, includeConvertedFromIssue, includeIntegrationsSettings, includeLead, includeCreator)
}

func ListUser(ctx context.Context, client graphql.Client, first int, after string, filter *UserFilter, includeOrganization *bool) (*listUserResponse, error) {
	return listUser(ctx, client, first, after, filter, includeOrganization)
}

func ListTeam(ctx context.Context, client graphql.Client, first int, after string, filter *TeamFilter, includeCycle *bool, includeIssueState *bool, includeTemplateForMembers *bool, includeTemplateForNonMembers *bool, includeWorkflowState *bool, includeIntegrationsSettings *bool, includeDuplicateWorkflowState *bool, includeOrganization *bool, includeReviewWorkflowState *bool, includeStartWorkflowState *bool, includeTriageWorkflowState *bool) (*listTeamResponse, error) {
	return listTeam(ctx, client, first, after, filter, includeCycle, includeIssueState, includeTemplateForMembers, includeTemplateForNonMembers, includeWorkflowState, includeIntegrationsSettings, includeDuplicateWorkflowState, includeOrganization, includeReviewWorkflowState, includeStartWorkflowState, includeTriageWorkflowState)
}

func ListWorkflowState(ctx context.Context, client graphql.Client, first int, after string, filter *WorkflowStateFilter, includeTeam *bool) (*listWorkflowStateResponse, error) {
	return listWorkflowState(ctx, client, first, after, filter, includeTeam)
}

func ListIssueLabel(ctx context.Context, client graphql.Client, first int, after string, filter *IssueLabelFilter, includeCreator *bool, includeOrganization *bool, includeParent *bool, includeTeam *bool) (*listIssueLabelResponse, error) {
	return listIssueLabel(ctx, client, first, after, filter, includeCreator, includeOrganization, includeParent, includeTeam)
}

func ListOrganization(ctx context.Context, client graphql.Client, includeSubscription *bool) (*listOrganizationResponse, error) {
	return listOrganization(ctx, client, includeSubscription)
}
