package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/google/go-querystring/query"
	"github.com/jeffail/gabs"
)

type SshKeyParams struct {
	Name      string `url:"name"`
	PublicKey string `url:"public_key"`
}

func SshKeysList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/sshkeys", HTTPGet, "")
}

func SshKeyCreate(params SshKeyParams) (json *gabs.Container, err error) {
	v, _ := query.Values(params)
	return makeJSONCall(config.URL()+"/v1/sshkeys", HTTPPost, v.Encode())
}

func SshKeyDelete(name string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/sshkeys/"+name, HTTPDelete, "")
}
