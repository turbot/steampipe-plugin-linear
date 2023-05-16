package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinearOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_organization",
		Description: "Linear Organization",
		List: &plugin.ListConfig{
			Hydrate: getOrganization,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the entity.",
			},
			{
				Name:        "allowed_auth_services",
				Type:        proto.ColumnType_JSON,
				Description: "Allowed authentication providers, empty array means all are allowed",
			},
			{
				Name:        "archived_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was archived. Null if the entity has not been archived.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was created.",
			},
			{
				Name:        "created_issue_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of issues in the organization.",
			},
			{
				Name:        "deletion_requested_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which deletion of the organization was requested.",
			},
			{
				Name:        "git_branch_format",
				Type:        proto.ColumnType_STRING,
				Description: "How git branches are formatted. If null, default formatting will be used.",
			},
			{
				Name:        "git_linkback_messages_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the Git integration linkback messages should be sent to private repositories.",
			},
			{
				Name:        "git_public_linkback_messages_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the Git integration linkback messages should be sent to public repositories.",
			},
			{
				Name:        "logo_url",
				Type:        proto.ColumnType_STRING,
				Description: "The organization's logo URL.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The organization's name.",
			},
			{
				Name:        "period_upload_volume",
				Type:        proto.ColumnType_DOUBLE,
				Description: "Rolling 30-day total upload volume for the organization, in megabytes.",
			},
			{
				Name:        "previous_url_keys",
				Type:        proto.ColumnType_JSON,
				Description: "Previously used URL keys for the organization (last 3 are kept and redirected).",
			},
			{
				Name:        "project_update_reminders_day",
				Type:        proto.ColumnType_STRING,
				Description: "The day at which to prompt for project updates.",
			},
			{
				Name:        "project_update_reminders_hour",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The hour at which to prompt for project updates.",
			},
			{
				Name:        "project_updates_reminder_frequency",
				Type:        proto.ColumnType_STRING,
				Description: "The frequency at which to prompt for project updates.",
			},
			{
				Name:        "release_channel",
				Type:        proto.ColumnType_STRING,
				Description: "The feature release channel the organization belongs to.",
			},
			{
				Name:        "roadmap_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the organization is using a roadmap.",
			},
			{
				Name:        "saml_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether SAML authentication is enabled for organization.",
			},
			{
				Name:        "scim_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether SCIM provisioning is enabled for organization.",
			},
			{
				Name:        "trial_ends_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the trial of the plus plan will end.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "url_key",
				Type:        proto.ColumnType_STRING,
				Description: "The organization's unique URL key.",
			},
			{
				Name:        "user_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of active users in the organization.",
			},
			{
				Name:        "subscription",
				Type:        proto.ColumnType_JSON,
				Description: "The organization's subscription to a paid plan.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The organization's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func getOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_organization.getOrganization", "connection_error", err)
		return nil, err
	}

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeSubscription := true
	if helpers.StringSliceContains(d.QueryContext.Columns, "subscription") {
		includeSubscription = false
	}

	getOrganizationResponse, err := gql.GetOrganization(ctx, conn, &includeSubscription)
	if err != nil {
		plugin.Logger(ctx).Error("linear_organization.getOrganization", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, getOrganizationResponse.Organization)

	return nil, nil
}
