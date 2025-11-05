# provider-terraform

This is a minimal Terraform provider named `provider-terraform` implemented in Go. It provides:

- a resource `automation_payload` with one attribute `json` (string)
- a data source `automation_payload` which accepts an `id` and returns `json` (computed)

Note: the provider stores payloads in an in-memory map inside the plugin process. This is for demo/testing only — state is not persisted across runs.

To build the provider binary:

```bash
cd provider-terraform
go mod tidy
go build -o terraform-provider-automation
```

To install the locally built provider:

1. Create the plugins directory structure:
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/hashicorp/automation/1.0.0/linux_amd64
```

2. Copy the provider binary to the plugins directory with the correct name:
```bash
cp terraform-provider-automation ~/.terraform.d/plugins/registry.terraform.io/hashicorp/automation/1.0.0/linux_amd64/terraform-provider-automation_v1.0.0
```

3. Make the provider binary executable:
```bash
chmod +x ~/.terraform.d/plugins/registry.terraform.io/hashicorp/automation/1.0.0/linux_amd64/terraform-provider-automation_v1.0.0
```

4. Update your Terraform configuration to use the local provider:
```hcl
terraform {
  required_providers {
    automation = {
      source = "registry.terraform.io/hashicorp/automation"
      version = "1.0.0"
    }
  }
}

resource "automation_payload" "example" {
  json = jsonencode({ message = "hello" })
}
```

5. Initialize Terraform to use the local provider:
```bash
terraform init
```

Note: When using a local provider, you may see a warning about incomplete lock file information. This is normal for local providers and won't affect functionality. If you need to use this provider on other platforms, you can generate additional checksums using:
```bash
terraform providers lock -platform=linux_amd64
```