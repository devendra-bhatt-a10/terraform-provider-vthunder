package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_SSL_CERT_REVOKE_RESOURCE = `
resource "vthunder_slb_ssl_cert_revoke" "SSLCertRevoke" {
	sampling_enable {
	    counters1 = "all"
	}
}
`

//Acceptance test
func TestAccSlbSSLCertRevoke_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_SSL_CERT_REVOKE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_ssl_cert_revoke.SSLCertRevoke", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
