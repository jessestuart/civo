package api

import (
	"fmt"

	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func AccountsList() (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/accounts", HTTPGet, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/accounts", HTTPGet, "")
	}
}

func AccountCreate(name string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/accounts", HTTPPost, fmt.Sprintf("name=%s", name))
	} else {
		return makeJSONCall(config.URL()+"/v1/accounts", HTTPPost, fmt.Sprintf("name=%s", name))
	}
}

func AccountDelete(name string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/accounts/"+name, HTTPDelete, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/accounts/"+name, HTTPDelete, "")
	}
}

func AccountReset(name string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/accounts/"+name, HTTPPut, fmt.Sprintf("name=%s", name))
	} else {
		return makeJSONCall(config.URL()+"/v1/accounts/"+name, HTTPPut, fmt.Sprintf("name=%s", name))
	}
}

func AccountFindByToken(token string) string {
	var accounts *gabs.Container
	var err error

	if Version() == 2 {
		accounts, err = makeJSONCall(config.URL()+"/v2/accounts", HTTPGet, "")
	} else {
		accounts, err = makeJSONCall(config.URL()+"/v1/accounts", HTTPGet, "")
	}
	if err != nil {
		fmt.Println(err)
		return ""
	}

	items, _ := accounts.Children()
	for _, child := range items {
		if child.S("api_key").Data().(string) == token {
			return child.S("username").Data().(string)
		}
	}
	return ""
}
