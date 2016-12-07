package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func RegionsList() (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/regions", HTTPGet, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/regions", HTTPGet, "")
	}
}
