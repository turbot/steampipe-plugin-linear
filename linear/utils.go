package linear

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

type linearClient struct {
	client   graphql.Client
	pageSize int64
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

func connect(ctx context.Context, d *plugin.QueryData) (*linearClient, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.(*linearClient), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	token := os.Getenv("LINEAR_TOKEN")

	// default Size
	pageSize := int64(50)

	linearConfig := GetConfig(d.Connection)
	if linearConfig.Token != nil {
		token = *linearConfig.Token
	}
	if linearConfig.PageSize != nil {
		// check if the provided value is more than the max limit
		if *linearConfig.PageSize > 250 {
			pageSize = 250
		} else {
			pageSize = *linearConfig.PageSize
		}
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	httpClient := http.Client{
		Transport: &authedTransport{
			key:     token,
			wrapped: http.DefaultTransport,
		},
	}
	graphqlClient := graphql.NewClient("https://api.linear.app/graphql", &httpClient)

	gqlClient := &linearClient{
		client:   graphqlClient,
		pageSize: pageSize,
	}

	return gqlClient, nil
}
