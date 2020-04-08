---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_templateDns"
sidebar_current: "docs-vthunder-resource-slb_templateDns"
description: |-
    Provides details about vthunder SLB template DNS resource for A10
---

# vthunder\_slb\_template\_dns

`vthunder_TemplateDNS` Provides details about vthunder SLB template DNS
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_dns" "templatedns" {
		name = "testdns"
        response_rate_limiting {
            enable_log=1
            filter_response_rate=3
            slip_rate=4
            response_rate=5
            window=6
            action="log-only"
        }
}
```

## Argument Reference

* `name` - DNS Template Name
* `enable_log` - Use private key for authentication
* `filter_response_rate` - Maximum allowed request rate for the filter. This should match average traffic. (default 20 per two seconds)
* `slip_rate` - Every n’th response that would be rate-limited will be let through instead
* `response_rate` - Responses exceeding this rate within the window will be dropped (default 5 per second)
* `window` - Rate-Limiting Interval in Seconds (default is one)
* `action` - 'log-only’: Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit’: Rate-Limit based on configuration (Default); 'whitelist’: Whitelist, disable rate-limiting;





