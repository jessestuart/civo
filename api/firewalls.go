package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/google/go-querystring/query"
	"github.com/jeffail/gabs"
)

type FirewallRuleParams struct {
	Protocol  string `url:"protocol"`
	StartPort string `url:"start_port"`
	EndPort   string `url:"end_port"`
	CIDR      string `url:"cidr"`
	Direction string `url:"direction"`
}

func FirewallsList() (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/firewalls", HTTPGet, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/firewalls", HTTPGet, "")
	}
}

func FirewallCreate(name string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/firewalls", HTTPPost, "name="+name)
	} else {
		return makeJSONCall(config.URL()+"/v1/firewalls", HTTPPost, "name="+name)
	}
}

func FirewallDestroy(name string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/firewalls/"+name, HTTPDelete, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/firewalls/"+name, HTTPDelete, "")
	}
}

func FirewallRules(name string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/firewalls/"+name+"/rules", HTTPGet, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/firewalls/"+name+"/rules", HTTPGet, "")
	}
}

func FirewallRuleCreate(name string, params FirewallRuleParams) (json *gabs.Container, err error) {
	v, _ := query.Values(params)
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/firewalls/"+name+"/rules", HTTPPost, v.Encode())
	} else {
		return makeJSONCall(config.URL()+"/v1/firewalls/"+name+"/rules", HTTPPost, v.Encode())
	}
}

func FirewallRuleDelete(name, id string) (json *gabs.Container, err error) {
	if Version() == 2 {
		return makeJSONCall(config.URL()+"/v2/firewalls/"+name+"/rules/"+id, HTTPDelete, "")
	} else {
		return makeJSONCall(config.URL()+"/v1/firewalls/"+name+"/rules/"+id, HTTPDelete, "")
	}
}
