# Copilot instructions for this repository

This repository contains a small Terraform provider plugin written in Go. The goal of these instructions
is to help an AI coding assistant be immediately productive when making changes or adding features.

- Big picture
  - This is a Terraform provider plugin (Go) whose binary is served from `main.go` via
    `plugin.Serve({ ProviderFunc: provider.Provider })`.
  - `provider/provider.go` constructs an in-process `PayloadStore` (an in-memory map protected
    by a `sync.RWMutex`) and registers resources and data sources.
  - Resource implementation lives in `provider/resource_payload.go` (CRUD functions that capture the
    store) and the data-source implementation is in `provider/data_source_payload.go`.
  - State is intentionally in-memory (not persisted). Changes will be lost when the plugin process
    exits. This is a deliberate design for demo/testing.

- Key files to inspect or update
  - `main.go` — plugin entrypoint (do not change how `plugin.Serve` is invoked unless you know
    Terraform plugin conventions).
  - `provider/provider.go` — register resources/data-sources and create the shared `PayloadStore`.
  - `provider/resource_payload.go` — resource schema and Create/Read/Update/Delete implementations.
  - `provider/data_source_payload.go` — data source schema and Read implementation.
  - `go.mod` — module and dependency versions (use `go mod tidy` to keep it clean).
  - `tests/` — contains example Terraform configurations and state for manual testing.

- Build, install, and manual test workflow (explicit)
  1. From the repository root (module root) run:
     ```bash
     go mod tidy
     go build -o terraform-provider-automation
     ```
  2. Install locally for Terraform testing (example for linux_amd64 and provider `hashicorp/automation`):
     ```bash
     mkdir -p ~/.terraform.d/plugins/registry.terraform.io/hashicorp/automation/1.0.0/linux_amd64
     cp terraform-provider-automation ~/.terraform.d/plugins/registry.terraform.io/hashicorp/automation/1.0.0/linux_amd64/terraform-provider-automation_v1.0.0
     chmod +x ~/.terraform.d/plugins/registry.terraform.io/hashicorp/automation/1.0.0/linux_amd64/terraform-provider-automation_v1.0.0
     terraform init
     ```

- Project-specific patterns and conventions
  - Resource/data functions accept a `store *PayloadStore` and return `schema.*ContextFunc` closures.
    Follow that approach when adding new resources or data sources so they share the same store instance.
  - Use `d.SetId(id)` and `d.Set("field", value)` consistently. For deleted resources set `d.SetId("")`.
  - Use `diag.FromErr(err)` to convert Go errors into Terraform diagnostics (see existing uses).
  - Concurrency: the store uses `RWMutex` — read-lock for reads and write-lock for writes/deletes.
  - IDs are generated with `github.com/google/uuid` (see `resource_payload.go`). Keep this pattern for
    stable, short-lived IDs unless a different ID scheme is required.

- Integration points and tests
  - The provider is invoked by Terraform as a plugin; unit tests should avoid starting the plugin.
    For small logic changes, prefer unit testing functions directly by referencing `PayloadStore`.
  - Manual integration: use the Terraform samples under `tests/` to exercise CRUD flows.

- When editing files, look for these common pitfalls
  - Do not assume persistence across runs — any new code that expects persistence must add it
    explicitly and update README/instructions.
  - Keep the Provider() signature stable; Terraform expects the provider factory to return `*schema.Provider`.
  - When returning errors from data sources, prefer `return diag.FromErr(fmt.Errorf(".."))` so Terraform shows a clear message.

If anything here is unclear or you'd like more examples (unit test snippets, how to add a second resource,
or a CI workflow to build and smoke-test the provider), tell me what to expand and I will iterate.
