package go_vthunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type SlbDNSResponseRateLimiting struct {
	UUID SlbDNSResponseRateLimitingInstance `json:"dns-response-rate-limiting,omitempty"`
}

type SamplingEnable17 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SlbDNSResponseRateLimitingInstance struct {
	Counters1 []SamplingEnable17 `json:"sampling-enable,omitempty"`
	UUID      string             `json:"uuid,omitempty"`
}

func GetSlbDNSResponseRateLimiting(id string, host string) (*SlbDNSResponseRateLimiting, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/dns-response-rate-limiting", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbDNSResponseRateLimiting
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}
}

func PostSlbDNSResponseRateLimiting(id string, vc SlbDNSResponseRateLimiting, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/dns-response-rate-limiting", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v SlbDNSResponseRateLimiting
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}
