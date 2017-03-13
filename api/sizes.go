package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func SizesList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/sizes", HTTPGet, "")
}
