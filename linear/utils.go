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

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

func connect(ctx context.Context, d *plugin.QueryData) (graphql.Client, error) {
	token := os.Getenv("LINEAR_TOKEN")

	linearConfig := GetConfig(d.Connection)
	if &linearConfig != nil {
		if linearConfig.Token != nil {
			token = *linearConfig.Token
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

	return graphqlClient, nil
}
