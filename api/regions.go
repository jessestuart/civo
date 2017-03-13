package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func RegionsList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/regions", HTTPGet, "")
}
