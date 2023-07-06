package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

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
				Description: "The unique identifier of the entity.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "allowed_auth_services",
				Description: "Allowed authentication providers, empty array means all are allowed",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "archived_at",
				Description: "The time at which the entity was archived. Null if the entity has not been archived.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_at",
				Description: "The time at which the entity was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_issue_count",
				Description: "Number of issues in the organization.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "deletion_requested_at",
				Description: "The time at which deletion of the organization was requested.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "git_branch_format",
				Description: "How git branches are formatted. If null, default formatting will be used.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "git_linkback_messages_enabled",
				Description: "Whether the Git integration linkback messages should be sent to private repositories.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "git_public_linkback_messages_enabled",
				Description: "Whether the Git integration linkback messages should be sent to public repositories.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "logo_url",
				Description: "The organization's logo URL.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The organization's name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "period_upload_volume",
				Description: "Rolling 30-day total upload volume for the organization, in megabytes.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "previous_url_keys",
				Description: "Previously used URL keys for the organization (last 3 are kept and redirected).",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "project_update_reminders_day",
				Description: "The day at which to prompt for project updates.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_update_reminders_hour",
				Description: "The hour at which to prompt for project updates.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "project_updates_reminder_frequency",
				Description: "The frequency at which to prompt for project updates.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "release_channel",
				Description: "The feature release channel the organization belongs to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "roadmap_enabled",
				Description: "Whether the organization is using a roadmap.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "saml_enabled",
				Description: "Whether SAML authentication is enabled for organization.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "scim_enabled",
				Description: "Whether SCIM provisioning is enabled for organization.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "trial_ends_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the trial of the plus plan will end.",
			},
			{
				Name:        "updated_at",
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "url_key",
				Description: "The organization's unique URL key.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_count",
				Description: "Number of active users in the organization.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "subscription",
				Description: "The organization's subscription to a paid plan.",
				Type:        proto.ColumnType_JSON,
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

// HYDRATE FUNCTION

func getOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_organization.getOrganization", "connection_error", err)
		return nil, err
	}

	getOrganizationResponse, err := gql.GetOrganization(ctx, conn.client)
	if err != nil {
		plugin.Logger(ctx).Error("linear_organization.getOrganization", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, getOrganizationResponse.Organization)

	return nil, nil
}
