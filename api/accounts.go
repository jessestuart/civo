package api

import (
	"fmt"

	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func AccountsList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/accounts", HTTPGet, "")
}

func AccountCreate(name string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/accounts", HTTPPost, fmt.Sprintf("name=%s", name))
}

func AccountDelete(name string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/accounts/"+name, HTTPDelete, "")
}
