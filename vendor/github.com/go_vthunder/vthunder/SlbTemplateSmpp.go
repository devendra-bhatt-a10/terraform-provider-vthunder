package go_vthunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type Smpp struct {
	Name SmppInstance `json:"smpp"`
}
type SmppInstance struct {
	Name                      string `json:"name"`
	ServerEnquireLink         int    `json:"server-enquire-link"`
	ServerSelectionPerRequest int    `json:"server-selection-per-request"`
	ClientEnquireLink         int    `json:"client-enquire-link"`
	UserTag                   string `json:"user-tag"`
	ServerEnquireLinkVal      int    `json:"server-enquire-link-val"`
	User                      string `json:"user"`
	Password                  string `json:"password"`
	UUID                      string `json:"uuid"`
}

func GetSmpp(id string, name string, host string) (*Smpp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/smpp/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Smpp
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

func PostSmpp(id string, sg Smpp, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/smpp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Smpp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func PutSmpp(id string, name string, sg Smpp, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/smpp/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Smpp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func DeleteSmpp(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/smpp/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Smpp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
