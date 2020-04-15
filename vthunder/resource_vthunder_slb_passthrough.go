package vthunder

//vThunder resource SlbPassthrough

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbPassthrough() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbPassthroughCreate,
		Update: resourceSlbPassthroughUpdate,
		Read:   resourceSlbPassthroughRead,
		Delete: resourceSlbPassthroughDelete,
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
		},
	}
}

func resourceSlbPassthroughCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbPassthrough (Inside resourceSlbPassthroughCreate) ")

		data := dataToSlbPassthrough(d)
		logger.Println("[INFO] received formatted data from method data to SlbPassthrough --")
		d.SetId("1")
		go_vthunder.PostSlbPassthrough(client.Token, data, client.Host)

		return resourceSlbPassthroughRead(d, meta)

	}
	return nil
}

func resourceSlbPassthroughRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbPassthrough (Inside resourceSlbPassthroughRead)")

	if client.Host != "" {
		name := d.Id()

		data, err := go_vthunder.GetSlbPassthrough(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}
func resourceSlbPassthroughUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbPassthroughRead(d, meta)
}

func resourceSlbPassthroughDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbPassthroughRead(d, meta)
}

func dataToSlbPassthrough(d *schema.ResourceData) go_vthunder.Passthrough {
	var vc go_vthunder.Passthrough
	var c go_vthunder.PassthroughInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable24, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable24
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
