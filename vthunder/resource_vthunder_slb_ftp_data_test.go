package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FTP_DATA_RESOURCE = `
resource "vthunder_slb_ftp_data" "ftp_data1" {
	sampling_enable {
		counters1 = "all"
	}
}

`

//Acceptance test
func TestAccVthunderSlbFTPData_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FTP_DATA_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_ftp_data.ftp_data1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
