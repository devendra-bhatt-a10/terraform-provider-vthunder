package vthunder

//vThunder resource SlbRateLimitLog

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbRateLimitLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbRateLimitLogCreate,
		Update: resourceSlbRateLimitLogUpdate,
		Read:   resourceSlbRateLimitLogRead,
		Delete: resourceSlbRateLimitLogDelete,
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

func resourceSlbRateLimitLogCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbRateLimitLog (Inside resourceSlbRateLimitLogCreate) ")

		data := dataToSlbRateLimitLog(d)
		logger.Println("[INFO] received formatted data from method data to SlbRateLimitLog --")
		d.SetId("1")
		go_vthunder.PostSlbRateLimitLog(client.Token, data, client.Host)

		return resourceSlbRateLimitLogRead(d, meta)

	}
	return nil
}

func resourceSlbRateLimitLogRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbRateLimitLog (Inside resourceSlbRateLimitLogRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbRateLimitLog(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbRateLimitLogUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbRateLimitLogRead(d, meta)
}

func resourceSlbRateLimitLogDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbRateLimitLogRead(d, meta)
}

func dataToSlbRateLimitLog(d *schema.ResourceData) go_vthunder.RateLimitLog {
	var vc go_vthunder.RateLimitLog
	var c go_vthunder.RateLimitLogInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable30, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable30
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
