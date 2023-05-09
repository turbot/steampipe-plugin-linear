package main

import (
	"github.com/steampipe-plugin-linear/linear"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: linear.Plugin})
}
