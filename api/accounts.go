package api

import (
	"fmt"

	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func AccountsList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/accounts", HTTPGet, "")
}

func AccountCreate(name string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/accounts", HTTPPost, fmt.Sprintf("name=%s", name))
}

func AccountDelete(name string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/accounts/"+name, HTTPDelete, "")
}

func AccountReset(name string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/accounts/"+name, HTTPPut, fmt.Sprintf("name=%s", name))
}

func AccountFindByAPIKey(apikey string) string {
	var accounts *gabs.Container
	var err error

	accounts, err = makeJSONCall(config.URL()+"/v2/accounts", HTTPGet, "")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	items, _ := accounts.Children()
	for _, child := range items {
		if child.S("api_key").Data().(string) == apikey {
			return child.S("id").Data().(string)
		}
	}
	return ""
}

func AccountFindByName(name string) string {
	var accounts *gabs.Container
	var err error

	accounts, err = makeJSONCall(config.URL()+"/v2/accounts", HTTPGet, "")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	items, _ := accounts.Children()
	for _, child := range items {
		if child.S("username").Data().(string) == name {
			return child.S("id").Data().(string)
		}
	}
	return ""
}
