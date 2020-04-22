package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type IcapHTTP struct {
	Counters1 IcapHTTPInstance `json:"icap_http,omitempty"`
}
type SamplingEnable40 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type IcapHTTPInstance struct {
	Counters1 []SamplingEnable40 `json:"sampling-enable,omitempty"`
}

func PostSlbIcapHTTP(id string, inst IcapHTTP, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbIcapHTTP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/icap_http", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IcapHTTP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetSlbIcapHTTP(id string, host string) (*IcapHTTP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbIcapHTTP")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/icap_http/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IcapHTTP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}
