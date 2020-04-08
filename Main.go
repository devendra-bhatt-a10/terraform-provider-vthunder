package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-vthunder/vthunder"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vthunder.Provider})
}
