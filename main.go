package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/stephanemetairie/provider-terraform/provider"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: provider.Provider,
    })
}
