package vthunder

// vThunder resource Slb FTPCtl

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbFTPCtl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbFTPCtlCreate,
		Update: resourceSlbFTPCtlUpdate,
		Read:   resourceSlbFTPCtlRead,
		Delete: resourceSlbFTPCtlDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceSlbFTPCtlCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating FTPCtl (Inside resourceSlbFTPCtlCreate)")

	if client.Host != "" {
		vc := dataToSlbFTPCtl(d)
		d.SetId("1")
		go_vthunder.PostSlbFTPCtl(client.Token, vc, client.Host)

		return resourceSlbFTPCtlRead(d, meta)
	}
	return nil
}

func resourceSlbFTPCtlRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading FTPCtl (Inside resourceSlbFTPCtlRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbFTPCtl(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No FTPCtl found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbFTPCtlUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFTPCtlRead(d, meta)
}

func resourceSlbFTPCtlDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFTPCtlRead(d, meta)
}

//Utility method to instantiate FTPCtl Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFTPCtl(d *schema.ResourceData) go_vthunder.FTPCtl {
	var vc go_vthunder.FTPCtl

	var c go_vthunder.FTPCtlInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable20, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable20
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
