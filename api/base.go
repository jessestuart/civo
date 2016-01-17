package api

import (
	"encoding/json"
	"fmt"

	"github.com/jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

var CurrentToken string

type HTTPMethod int

const (
	HTTPGet HTTPMethod = iota
	HTTPPost
	HTTPPut
	HTTPDelete
)

func Connect(token string) {
	CurrentToken = token
}

func requestHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("bearer %s", CurrentToken)
	headers["User-agent"] = "Civo Go client v1.0"
	return headers
}

func makeJSONCall(url string, method HTTPMethod, data string) (*gabs.Container, error) {
	request := gorequest.New()

	switch method {
	case HTTPGet:
		request = request.Get(url)
	case HTTPPost:
		request = request.Post(url).Send(data)
	case HTTPPut:
		request = request.Put(url).Send(data)
	case HTTPDelete:
		request = request.Delete(url)
	}

	for name, value := range requestHeaders() {
		request = request.Set(name, value)
	}
	resp, body, errs := request.End()
	if errs != nil {
		return nil, errs[0]
	}
	status := resp.StatusCode

	if !(status >= 200 && status <= 299) {
		return nil, HTTPErrorNew(fmt.Sprintf("Unable to make Openstack API call: %s", body), url, status)
	}
	if body != "" {
		jsonObject, err := gabs.ParseJSON([]byte(body))
		if err != nil {
			var object interface{}
			jsonObject := gabs.New()
			_ = json.Unmarshal([]byte([]byte(body)), &object)
			jsonObject.SetP(object, "items")
			return jsonObject, nil
		}
		return jsonObject, nil
	}
	return nil, nil
}
