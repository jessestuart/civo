package api

import (
	"fmt"
	"strings"

	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func NetworksList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/networks", HTTPGet, "")
}

func NetworkCreate(name, label, region string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/networks", HTTPPost, "name="+name+"&label="+label+"&region="+region)
}

func NetworkDestroy(id string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v2/networks/"+id, HTTPDelete, "")
}

// Utility functions ---------------------------------------------------------------------------------------------------

func NetworkFind(search string) string {
	instances, err := NetworksList()
	if err != nil {
		fmt.Println("DEBUG: Returning early because err is", err)
		return ""
	}
	items, _ := instances.Children()
	for _, child := range items {
		id := child.S("id").Data().(string)
		label := child.S("label").Data().(string)
		if strings.Contains(id, search) {
			return id
		}
		if strings.Contains(label, search) {
			return id
		}
	}
	return ""
}
