---
page_title: "cloudflare_rulesets Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_rulesets (Data Source)



## Example Usage

```terraform
data "cloudflare_rulesets" "example_rulesets" {
  account_id = "account_id"
  zone_id = "zone_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `max_items` (Number) Max items to fetch, default: 1000
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `description` (String) An informative description of the ruleset.
- `id` (String) The unique ID of the ruleset.
- `kind` (String) The kind of the ruleset.
Available values: "managed", "custom", "root", "zone".
- `name` (String) The human-readable name of the ruleset.
- `phase` (String) The phase of the ruleset.
Available values: "ddos_l4", "ddos_l7", "http_config_settings", "http_custom_errors", "http_log_custom_fields", "http_ratelimit", "http_request_cache_settings", "http_request_dynamic_redirect", "http_request_firewall_custom", "http_request_firewall_managed", "http_request_late_transform", "http_request_origin", "http_request_redirect", "http_request_sanitize", "http_request_sbfm", "http_request_transform", "http_response_compression", "http_response_firewall_managed", "http_response_headers_transform", "magic_transit", "magic_transit_ids_managed", "magic_transit_managed", "magic_transit_ratelimit".


