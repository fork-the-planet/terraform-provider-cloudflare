---
page_title: "cloudflare_cloud_connector_rules Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_cloud_connector_rules (Resource)



## Example Usage

```terraform
resource "cloudflare_cloud_connector_rules" "example_cloud_connector_rules" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  rules = [{
    id = "95c365e17e1b46599cd99e5b231fac4e"
    description = "Rule description"
    enabled = true
    expression = "http.cookie eq \"a=b\""
    parameters = {
      host = "examplebucket.s3.eu-north-1.amazonaws.com"
    }
    provider = "aws_s3"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier.

### Optional

- `rules` (Attributes List) (see [below for nested schema](#nestedatt--rules))

### Read-Only

- `id` (String) Identifier.

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `description` (String)
- `enabled` (Boolean)
- `expression` (String)
- `parameters` (Attributes) Parameters of Cloud Connector Rule (see [below for nested schema](#nestedatt--rules--parameters))
- `provider` (String) Cloud Provider type
Available values: "aws_s3", "cloudflare_r2", "gcp_storage", "azure_storage".

Read-Only:

- `id` (String)

<a id="nestedatt--rules--parameters"></a>
### Nested Schema for `rules.parameters`

Optional:

- `host` (String) Host to perform Cloud Connection to

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_cloud_connector_rules.example '<zone_id>'
```
