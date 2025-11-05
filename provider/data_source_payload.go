package provider

import (
    "context"
    "fmt"

    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePayload(store *PayloadStore) *schema.Resource {
    return &schema.Resource{
        ReadContext: dataSourcePayloadRead(store),
        Schema: map[string]*schema.Schema{
            "id": {
                Type:     schema.TypeString,
                Required: true,
            },
            "json": {
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}

func dataSourcePayloadRead(store *PayloadStore) schema.ReadContextFunc {
    return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        var diags diag.Diagnostics

        id := d.Get("id").(string)

        store.RLock()
        jsonVal, ok := store.data[id]
        store.RUnlock()

        if !ok {
            return diag.FromErr(fmt.Errorf("payload not found: %s", id))
        }

        d.SetId(id)
        if err := d.Set("json", jsonVal); err != nil {
            return diag.FromErr(err)
        }

        return diags
    }
}
