package linear

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type linearConfig struct {
	Token    *string `hcl:"token"`
	PageSize *int64  `hcl:"page_size"`
}

func ConfigInstance() interface{} {
	return &linearConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) linearConfig {
	if connection == nil || connection.GetConfig() == nil {
		return linearConfig{}
	}
	config, _ := connection.GetConfig().(linearConfig)
	return config
}
