package main

import (
	"github.com/bluesentry/terraform-provider-appstream/appstream"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: appstream.Provider})
}
