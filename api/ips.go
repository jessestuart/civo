package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func IPCreate(instanceID string, public bool) (json *gabs.Container, err error) {
	params := ""
	if public {
		params = "public=true"
	}
	return makeJSONCall(config.URL()+"/v2/instances/"+instanceID+"/ip", HTTPPost, params)
}

func IPDelete(instanceID, ip string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/instances/"+instanceID+"/ip/"+ip, HTTPDelete, "")
}

func IPConnect(instanceID, publicIP, privateIP string) (json *gabs.Container, err error) {
	params := "private_ip=" + privateIP
	return makeJSONCall(config.URL()+"/v2/instances/"+instanceID+"/ip"+publicIP, HTTPPut, params)
}
