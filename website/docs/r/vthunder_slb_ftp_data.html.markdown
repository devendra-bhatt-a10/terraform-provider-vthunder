---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_ftp_data"
sidebar_current: "docs-vthunder-resource-slb-ftp-data"
description: |-
    Provides details about vthunder SLB ftp-data resource for A10
---

# vthunder\_slb\_ftp_data

`vthunder_slb_ftp_data` Provides details about vthunder SLB ftp-data
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_ftp_data" "ftp_data" {
	sampling_enable {
		counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable` - 
	* `counters1` - 'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;


