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
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/instances/"+instanceID+"/ip", HTTPPost, params)
	} else {
		return makeJSONCall(config.URL()+"/v1/instances/"+instanceID+"/ip", HTTPPost, params)
	}
}

func IPDelete(instanceID, ip string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/instances/"+instanceID+"/ip/"+ip, HTTPDelete, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/instances/"+instanceID+"/ip/"+ip, HTTPDelete, "")
	}
}

func IPConnect(instanceID, publicIP, privateIP string) (json *gabs.Container, err error) {
	params := "private_ip=" + privateIP
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/instances/"+instanceID+"/ip"+publicIP, HTTPPut, params)
	} else {
		return makeJSONCall(config.URL()+"/v1/instances/"+instanceID+"/ip"+publicIP, HTTPPut, params)
	}
}
