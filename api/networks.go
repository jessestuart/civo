package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func NetworksList() (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/networks", HTTPGet, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/networks", HTTPGet, "")
	}
}

func NetworkCreate(name, label string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/networks", HTTPPost, "name="+name+"&label="+label)
	} else {
		return makeJSONCall(config.URL()+"/v1/networks", HTTPPost, "name="+name+"&label="+label)
	}
}

func NetworkDestroy(id string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/networks/"+id, HTTPDelete, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/networks/"+id, HTTPDelete, "")
	}
}
