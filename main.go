package main

import (
	"github.com/allence-tunisie/terraform-provider-atn/atn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: atn.Provider,
	})
}
