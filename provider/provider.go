package provider

import (
    "sync"

    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// PayloadStore is a simple in-memory store for payloads.
type PayloadStore struct {
    sync.RWMutex
    data map[string]string
}

// Provider returns the terraform provider schema, resources and data sources.
func Provider() *schema.Provider {
    store := &PayloadStore{
        data: make(map[string]string),
    }

    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
            // Resource type is automation_payload
            "automation_payload": resourcePayload(store),
        },
        DataSourcesMap: map[string]*schema.Resource{
            // Data source is automation_payload
            "automation_payload": dataSourcePayload(store),
        },
    }
}
