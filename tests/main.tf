terraform {
  required_providers {
    automation = {
      source  = "registry.terraform.io/hashicorp/automation"
      version = "1.0.0"
    }
  }
}


locals {
  json_content = file("${path.module}/ba-test-1.json")
}

resource "automation_payload" "complex_example" {
  json = local.json_content
}

# output "complex_payload_id" {
#   value = automation_payload.complex_example.id
# }

# data "automation_payload" "complex_read" {
#   id = automation_payload.complex_example.id
# }

# output "complex_payload_json" {
#   value = data.automation_payload.complex_read.json
# }