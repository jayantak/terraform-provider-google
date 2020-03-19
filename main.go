package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/jayantak/terraform-provider-google/google"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider})
}
