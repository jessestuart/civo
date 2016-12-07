package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func SizesList() (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/sizes", HTTPGet, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/sizes", HTTPGet, "")
	}
}
