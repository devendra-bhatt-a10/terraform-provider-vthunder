package vthunder

// vThunder resource Slb Aflow

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbAflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbAflowCreate,
		Update: resourceSlbAflowUpdate,
		Read:   resourceSlbAflowRead,
		Delete: resourceSlbAflowDelete,

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

func resourceSlbAflowCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating aflow (Inside resourceSlbAflowCreate)")

	if client.Host != "" {
		vc := dataToSlbAflow(d)
		d.SetId("1")
		go_vthunder.PostSlbAflow(client.Token, vc, client.Host)

		return resourceSlbAflowRead(d, meta)
	}
	return nil
}

func resourceSlbAflowRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading aflow (Inside resourceSlbAflowRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbAflow(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No aflow found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbAflowUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbAflowRead(d, meta)
}

func resourceSlbAflowDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbAflowRead(d, meta)
}

//Utility method to instantiate Aflow Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbAflow(d *schema.ResourceData) go_vthunder.Aflow {
	var vc go_vthunder.Aflow

	var c go_vthunder.AflowInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable6, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable6
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
