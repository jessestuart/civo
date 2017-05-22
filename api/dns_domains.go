package api

import (
	"fmt"
	"strings"

	"github.com/absolutedevops/civo/config"
	"github.com/google/go-querystring/query"
	"github.com/jeffail/gabs"
)

type DnsRecordParams struct {
	Name     string `url:"name"`
	Value    string `url:"value"`
	Type     string `url:"type"`
	Priority string `url:"priority"`
	TTL      string `url:"ttl"`
}

func DnsDomainsList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/dns", HTTPGet, "")
}

func DnsDomainCreate(name string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/dns", HTTPPost, "name="+name)
}

func DnsDomainDestroy(id string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/dns/"+id, HTTPDelete, "")
}

func DnsRecords(id string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/dns/"+id+"/records", HTTPGet, "")
}

func DnsRecordCreate(id string, params DnsRecordParams) (json *gabs.Container, err error) {
	v, _ := query.Values(params)
	return makeJSONCall(config.URL()+"/v2/dns/"+id+"/records", HTTPPost, v.Encode())
}

func DnsRecordDelete(id, recordID string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/dns/"+id+"/records/"+recordID, HTTPDelete, "")
}

// Utility functions ---------------------------------------------------------------------------------------------------

func DnsDomainFind(search string) string {
	domains, err := DnsDomainsList()
	if err != nil {
		fmt.Println("DEBUG: Returning early because err is", err)
		return ""
	}
	items, _ := domains.S("items").Children()
	for _, child := range items {
		id := child.S("id").Data().(string)
		name := child.S("name").Data().(string)
		if strings.Contains(id, search) {
			return id
		}
		if strings.Contains(name, search) {
			return id
		}
	}
	return ""
}
