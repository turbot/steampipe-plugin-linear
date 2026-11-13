package main

import (
	"github.com/turbot/steampipe-plugin-linear/linear"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: linear.Plugin})
}
