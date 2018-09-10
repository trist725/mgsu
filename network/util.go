package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpRequestResponse(
	method string,
	url string,
	timeout time.Duration,
	token string,
	request interface{},
	response interface{},
) error {

	var req *http.Request

	if request != nil {
		requestJSON, err := json.Marshal(request)
		if err != nil {
			return fmt.Errorf("marshal requestBody[%#v] to json fail, %s", request, err)
		}
		req, err = http.NewRequest(method, fmt.Sprintf("http://%s", url), bytes.NewReader(requestJSON))
		if err != nil {
			return fmt.Errorf("new http request fail, %s", err)
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		var err error
		req, err = http.NewRequest(method, fmt.Sprintf("http://%s", url), nil)
		if err != nil {
			return fmt.Errorf("new http request fail, %s", err)
		}
	}

	if token != "" {
		req.Header.Add("Authorization", token)
	}

	client := &http.Client{Timeout: timeout}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do http request fail, %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("get data from res.Body fail")
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("unmarshal response[%T] from [%s] fail, %s", response, body, err)
	}

	return nil
}

func HttpGetJson(
	url string,
	timeout time.Duration,
	token string,
	request interface{},
	response interface{},
) error {
	return HttpRequestResponse("GET", url, timeout, token, request, response)
}

func HttpPostJson(
	url string,
	timeout time.Duration,
	token string,
	request interface{},
	response interface{},
) error {
	return HttpRequestResponse("POST", url, timeout, token, request, response)
}

func CollectHttpRequestParams(r *http.Request) (params map[string]string) {
	params = make(map[string]string)
	for k, v := range r.Form {
		params[k] = v[0]
	}
	return
}
