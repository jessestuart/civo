package api

import (
	"fmt"
	"strings"

	"github.com/absolutedevops/civo/config"
	"github.com/google/go-querystring/query"
	"github.com/jeffail/gabs"
)

type SshKeyParams struct {
	Name      string `url:"name"`
	PublicKey string `url:"public_key"`
}

func SshKeysList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/sshkeys", HTTPGet, "")
}

func SshKeyCreate(params SshKeyParams) (json *gabs.Container, err error) {
	v, _ := query.Values(params)
	return makeJSONCall(config.URL()+"/v2/sshkeys", HTTPPost, v.Encode())
}

func SshKeyDelete(name string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/sshkeys/"+name, HTTPDelete, "")
}

// Utility functions ---------------------------------------------------------------------------------------------------

func SshKeyFind(search string) string {
	instances, err := SshKeysList()
	if err != nil {
		fmt.Println("DEBUG: Returning early because err is", err)
		return ""
	}
	items, _ := instances.Children()
	for _, child := range items {
		id := child.S("id").Data().(string)
		label := child.S("name").Data().(string)
		if strings.Contains(id, search) {
			return id
		}
		if strings.Contains(label, search) {
			return id
		}
	}
	return ""
}
