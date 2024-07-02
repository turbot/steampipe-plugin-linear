package linear

import (
	"context"

	"github.com/turbot/steampipe-plugin-linear/gql"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "organization_id",
			Description: "Unique identifier for the organization.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getOrganizationId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getOrganizationIdMemoized = plugin.HydrateFunc(getOrganizationIdUncached).Memoize(memoize.WithCacheKeyFunction(getOrganizationIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getOrganizationId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getOrganizationIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getOrganizationIdCacheKey.
func getOrganizationIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getOrganizationId"
	return key, nil
}

func getOrganizationIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationInfoUncached", "connection_error", err)
		return nil, err
	}

	getOrganizationResponse, err := gql.GetOrganization(ctx, conn.client)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationInfoUncached", "api_error", err)
		return nil, err
	}

	return getOrganizationResponse.GetOrganization().Id, nil
}
