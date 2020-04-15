package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FTP_CTL_RESOURCE = `
resource "vthunder_slb_ftp_ctl" "ftp_ctl1" {
	sampling_enable {
		counters1 = "all"
	}
}

`

//Acceptance test
func TestAccVthunderFTPCtl_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FTP_CTL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_ftp_ctl.ftp_ctl1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
