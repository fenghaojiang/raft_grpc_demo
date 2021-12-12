package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const registerCenterAddr = "127.0.0.1:60000"

func DoGet(key string) string {
	resp, err := http.Get(fmt.Sprintf("http://%s/key/%s", registerCenterAddr, key))
	if err != nil {
		fmt.Println("failed to GET key:", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read response:", err.Error())
	}
	return string(body)
}

func DoSet(key, value string) {
	b, err := json.Marshal(map[string]string{key: value})
	if err != nil {
		fmt.Println("failed to encode key and value for POST:", err.Error())
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/key", registerCenterAddr), "application-type/json", bytes.NewReader(b))
	if err != nil {
		fmt.Println("POST request failed: ", err.Error())
	}
	defer resp.Body.Close()
}

func DoDelete(key string) {
	ru, err := url.Parse(fmt.Sprintf("http://%s/key/%s", registerCenterAddr, key))
	if err != nil {
		fmt.Println("failed to parse URL for delete:", err.Error())
	}
	req := &http.Request{
		Method: "DELETE",
		URL:    ru,
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("failed to GET key:", err.Error())
	}
	defer resp.Body.Close()
}
