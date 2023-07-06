package linear

import (
	"context"

	"github.com/steampipe-plugin-linear/gql"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableLinearUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linear_user",
		Description: "Linear User",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
			KeyColumns: []*plugin.KeyColumn{
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
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getUser,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier of the entity.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "active",
				Description: "Whether the user account is active or disabled (suspended).",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "admin",
				Description: "Whether the user is an organization administrator.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "archived_at",
				Description: "The time at which the entity was archived. Null if the entity has not been archived.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "avatar_url",
				Description: "An URL to the user's avatar image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time at which the entity was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "created_issue_count",
				Description: "Number of issues created.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "description",
				Description: "A short description of the user, either its title or bio.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "disable_reason",
				Description: "Reason why is the account disabled.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_name",
				Description: "The user's display (nick) name. Unique within each organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The user's email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "guest",
				Description: "Whether the user is a guest in the workspace and limited to accessing a subset of teams.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "invite_hash",
				Description: "Unique hash for the user to be used in invite URLs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_me",
				Description: "Whether the user is the currently authenticated user.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_seen",
				Description: "The last time the user was seen online. If null, the user is currently online.",
				Type:        proto.ColumnType_TIMESTAMP,

			},
			{
				Name:        "name",
				Description: "The user's full name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status_emoji",
				Description: "The emoji to represent the user current status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status_label",
				Description: "The label of the user current status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status_until_at",
				Description: "A date at which the user current status should be cleared.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "timezone",
				Description: "The local timezone of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "The last time at which the entity was meaningfully updated, i.e. for all changes of syncable properties except those for which updates should not produce an update to updatedAt (see skipUpdatedAtKeys). This is the same as the creation time if the entity hasn't been updated after creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "url",
				Description: "User's profile URL.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization",
				Description: "Organization the user belongs to.",
				Type:        proto.ColumnType_JSON,
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

// LIST FUNCTION

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_user.listUsers", "connection_error", err)
		return nil, err
	}

	var endCursor string

	// set page size
	var pageSize int = int(conn.pageSize)
	if d.QueryContext.Limit != nil {
		if int(*d.QueryContext.Limit) < pageSize {
			pageSize = int(*d.QueryContext.Limit)
		}
	}

	// set the requested filters
	filters := setUserFilters(d, ctx)

	for {
		listUserResponse, err := gql.ListUsers(ctx, conn.client, pageSize, endCursor, true, &filters)
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

// HYDRATE FUNCTION

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linear_user.getUser", "connection_error", err)
		return nil, err
	}

	getUserResponse, err := gql.GetUser(ctx, conn.client, &id)
	if err != nil {
		plugin.Logger(ctx).Error("linear_user.getUser", "api_error", err)
		return nil, err
	}

	return getUserResponse.User, nil
}

// Set the requested filter
func setUserFilters(d *plugin.QueryData, ctx context.Context) gql.UserFilter {
	var filter gql.UserFilter
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
