package linear

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type linearConfig struct {
	Token    *string `cty:"token"`
	PageSize *int64  `cty:"page_size"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"page_size": {
		Type: schema.TypeInt,
	},
}

func ConfigInstance() interface{} {
	return &linearConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) linearConfig {
	if connection == nil || connection.Config == nil {
		return linearConfig{}
	}
	config, _ := connection.Config.(linearConfig)
	return config
}
