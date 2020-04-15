package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_CONNECTION_REUSE_RESOURCE = `
resource "vthunder_slb_connection_reuse" "connection_reuse" {
	sampling_enable {
		counters1 = "all"
	}
}

`

//Acceptance test
func TestAccVthunderConnectionReuse_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_CONNECTION_REUSE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_connection_reuse.connection_reuse", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
