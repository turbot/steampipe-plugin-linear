package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinearUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_user",
		Description: "Linear User",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "id",
					Require: plugin.Optional,
				},
				{
					Name:    "active",
					Require: plugin.Optional,
				},
				{
					Name:    "admin",
					Require: plugin.Optional,
				},
				{
					Name:    "display_name",
					Require: plugin.Optional,
				},
				{
					Name:      "created_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:      "updated_at",
					Require:   plugin.Optional,
					Operators: []string{"=", ">", ">=", "<=", "<"},
				},
				{
					Name:    "email",
					Require: plugin.Optional,
				},
				{
					Name:    "is_me",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier of the entity.",
			},
			{
				Name:        "active",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the user account is active or disabled (suspended).",
			},
			{
				Name:        "admin",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the user is an organization administrator.",
			},
			{
				Name:        "archived_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was archived. Null if the entity has not been archived.",
			},
			{
				Name:        "avatar_url",
				Type:        proto.ColumnType_STRING,
				Description: "An URL to the user's avatar image.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the entity was created.",
			},
			{
				Name:        "created_issue_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of issues created.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "A short description of the user, either its title or bio.",
			},
			{
				Name:        "disable_reason",
				Type:        proto.ColumnType_STRING,
				Description: "Reason why is the account disabled.",
			},
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "The user's display (nick) name. Unique within each organization.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The user's email address.",
			},
			{
				Name:        "guest",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the user is a guest in the workspace and limited to accessing a subset of teams.",
			},
			{
				Name:        "invite_hash",
				Type:        proto.ColumnType_STRING,
				Description: "Unique hash for the user to be used in invite URLs.",
			},
			{
				Name:        "is_me",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the user is the currently authenticated user.",
			},
			{
				Name:        "last_seen",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time the user was seen online. If null, the user is currently online.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The user's full name.",
			},
			{
				Name:        "status_emoji",
				Type:        proto.ColumnType_STRING,
				Description: "The emoji to represent the user current status.",
			},
			{
				Name:        "status_label",
				Type:        proto.ColumnType_STRING,
				Description: "The label of the user current status.",
			},
			{
				Name:        "status_until_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "A date at which the user current status should be cleared.",
			},
			{
				Name:        "timezone",
				Type:        proto.ColumnType_STRING,
				Description: "The local timezone of the user.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "User's profile URL.",
			},
			{
				Name:        "organization",
				Type:        proto.ColumnType_JSON,
				Description: "Organization the user belongs to.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "The user's title.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_user.listUsers", "connection_error", err)
		return nil, err
	}

	var endCursor string
	var pageSize int = 50
	if d.QueryContext.Limit != nil {
		if int(*d.QueryContext.Limit) < pageSize {
			pageSize = int(*d.QueryContext.Limit)
		}
	}

	// By default, nested objects are excluded, and they will only be included if they are requested.
	includeOrganization := true
	if helpers.StringSliceContains(d.QueryContext.Columns, "organization") {
		includeOrganization = false
	}

	// set the requested filters
	filters := setUserFilters(d, ctx)

	for {
		listUserResponse, err := gql.ListUser(ctx, conn, pageSize, endCursor, true, &filters, &includeOrganization)
		if err != nil {
			plugin.Logger(ctx).Error("linear_user.listUsers", "api_error", err)
			return nil, err
		}

		for _, node := range listUserResponse.Users.Nodes {
			d.StreamListItem(ctx, node)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if !*listUserResponse.Users.PageInfo.HasNextPage {
			break
		}
		endCursor = *listUserResponse.Users.PageInfo.EndCursor
	}

	return nil, nil
}

// Set the requested filter
func setUserFilters(d *plugin.QueryData, ctx context.Context) gql.UserFilter {
	var filter gql.UserFilter
	if d.EqualsQuals["id"] != nil {
		id := &gql.IDComparator{
			Eq: types.String(d.EqualsQualString("id")),
		}
		filter.Id = id
	}
	if d.Quals["created_at"] != nil {
		createdAt := &gql.DateComparator{}
		for _, q := range d.Quals["created_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				createdAt.Eq = timestamp
			case ">":
				createdAt.Gt = timestamp
			case ">=":
				createdAt.Gte = timestamp
			case "<":
				createdAt.Lt = timestamp
			case "<=":
				createdAt.Lte = timestamp
			}
		}
		filter.CreatedAt = createdAt
	}
	if d.Quals["updated_at"] != nil {
		updatedAt := &gql.DateComparator{}
		for _, q := range d.Quals["updated_at"].Quals {
			timestamp := types.Time(q.Value.GetTimestampValue().AsTime())
			switch q.Operator {
			case "=":
				updatedAt.Eq = timestamp
			case ">":
				updatedAt.Gt = timestamp
			case ">=":
				updatedAt.Gte = timestamp
			case "<":
				updatedAt.Lt = timestamp
			case "<=":
				updatedAt.Lte = timestamp
			}
		}
		filter.UpdatedAt = updatedAt
	}
	if d.EqualsQuals["name"] != nil {
		name := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("name")),
		}
		filter.Name = name
	}
	if d.EqualsQuals["display_name"] != nil {
		displayName := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("display_name")),
		}
		filter.DisplayName = displayName
	}
	if d.EqualsQuals["email"] != nil {
		email := &gql.StringComparator{
			Eq: types.String(d.EqualsQualString("email")),
		}
		filter.Email = email
	}
	if d.EqualsQuals["active"] != nil {
		active := &gql.BooleanComparator{
			Eq: types.Bool(d.EqualsQuals["active"].GetBoolValue()),
		}
		filter.Active = active
	}
	if d.EqualsQuals["admin"] != nil {
		admin := &gql.BooleanComparator{
			Eq: types.Bool(d.EqualsQuals["admin"].GetBoolValue()),
		}
		filter.Admin = admin
	}
	if d.EqualsQuals["is_me"] != nil {
		isMe := &gql.BooleanComparator{
			Eq: types.Bool(d.EqualsQuals["is_me"].GetBoolValue()),
		}
		filter.IsMe = isMe
	}

	return filter
}
