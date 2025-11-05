package provider

import (
    "context"

    "github.com/google/uuid"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePayload(store *PayloadStore) *schema.Resource {
    return &schema.Resource{
        CreateContext: resourcePayloadCreate(store),
        ReadContext:   resourcePayloadRead(store),
        UpdateContext: resourcePayloadUpdate(store),
        DeleteContext: resourcePayloadDelete(store),

        Schema: map[string]*schema.Schema{
            "json": {
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourcePayloadCreate(store *PayloadStore) schema.CreateContextFunc {
    return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        var diags diag.Diagnostics

        jsonVal := d.Get("json").(string)
        id := uuid.New().String()

        store.Lock()
        store.data[id] = jsonVal
        store.Unlock()

        d.SetId(id)
        // also set back the json (in case provider normalizes it later)
        if err := d.Set("json", jsonVal); err != nil {
            return diag.FromErr(err)
        }

        return diags
    }
}

func resourcePayloadRead(store *PayloadStore) schema.ReadContextFunc {
    return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        var diags diag.Diagnostics

        id := d.Id()
        store.RLock()
        jsonVal, ok := store.data[id]
        store.RUnlock()

        if !ok {
            // resource not found
            d.SetId("")
            return diags
        }

        if err := d.Set("json", jsonVal); err != nil {
            return diag.FromErr(err)
        }
        return diags
    }
}

func resourcePayloadUpdate(store *PayloadStore) schema.UpdateContextFunc {
    return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        var diags diag.Diagnostics
        id := d.Id()
        jsonVal := d.Get("json").(string)

        store.Lock()
        store.data[id] = jsonVal
        store.Unlock()

        return diags
    }
}

func resourcePayloadDelete(store *PayloadStore) schema.DeleteContextFunc {
    return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
        var diags diag.Diagnostics
        id := d.Id()

        store.Lock()
        delete(store.data, id)
        store.Unlock()

        d.SetId("")
        return diags
    }
}
