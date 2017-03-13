package api

import (
	"errors"
	"fmt"

	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

var CurrentToken string

type HTTPMethod int

var discoveredVersion float64

const (
	HTTPGet HTTPMethod = iota
	HTTPPost
	HTTPPut
	HTTPDelete
)

func requestHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("bearer %s", config.CurrentToken())
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

	fmt.Println(">>> ", url)
	fmt.Println(body)

	if !(status >= 200 && status <= 299) {
		jsonObject, err := gabs.ParseJSON([]byte(body))
		if err != nil {
			return nil, err
		}
		return nil, errors.New(jsonObject.Path("reason").Data().(string))
	}
	if body != "" {
		jsonObject, err := gabs.ParseJSON([]byte(body))
		if err != nil {
			return nil, err
		}
		return jsonObject, nil
	}
	return nil, nil
}

func Version() float64 {
	if discoveredVersion == 0 {
		result, _ := makeJSONCall(config.URL()+"/ping", HTTPGet, "")
		discoveredVersion = result.S("version").Data().(float64)
	}

	return discoveredVersion
}
