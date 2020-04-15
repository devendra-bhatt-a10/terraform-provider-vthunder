---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_server_ssl"
sidebar_current: "docs-vthunder-resource-slb-template-server-ssl"
description: |-
    Provides details about vthunder slb template server-ssl resource for A10
---

# vthunder\_slb\_template\_server\_ssl

`vthunder_slb_template_server_ssl` provides details about slb template server_ssl
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_server_ssl" "server_ssl" {
	name = "testserverssl"
	user_tag = "test_tag"
	sslilogging = "disable"
	ocsp_stapling = 1
	ssli_logging = 1
	session_cache_size = 1
	handshake_logging_enable = 1
	close_notify = 1
}
```

## Argument Reference

* `name` - Server SSL Template Name
* `user_tag` - Customized tag
* `sslilogging` - 'disable’: Disable all logging; 'all’: enable all logging(error, info);
* `ocsp_stapling` - Enable ocsp-stapling support
* `ssli_logging` - SSLi logging level, default is error logging only
* `session_cache_size` - Session Cache Size (Maximum cache size. Default value 0 (Session ID reuse disabled))
* `handshake_logging_enable` - Enable SSL handshake logging
* `close_notify` - Send close notification when terminate connection

