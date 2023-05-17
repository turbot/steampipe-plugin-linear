package linear

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-linear",
		DefaultTransform: transform.FromCamel(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"linear_attachment":      tableLinearAttachment(ctx),
			"linear_comment":         tableLinearComment(ctx),
			"linear_integration":     tableLinearIntegration(ctx),
			"linear_issue":           tableLinearIssue(ctx),
			"linear_issue_label":     tableLinearIssueLabel(ctx),
			"linear_organization":    tableLinearOrganization(ctx),
			"linear_project":         tableLinearProject(ctx),
			"linear_team":            tableLinearTeam(ctx),
			"linear_team_membership": tableLinearTeamMembership(ctx),
			"linear_user":            tableLinearUser(ctx),
		},
	}
	return p
}
