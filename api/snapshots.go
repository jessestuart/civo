package api

import (
	"fmt"
	"strings"

	"github.com/absolutedevops/civo/config"
	"github.com/jeffail/gabs"
)

func SnapshotsList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/snapshots", HTTPGet, "")
}

func SnapshotCreate(name, instance_id string, safe bool) (json *gabs.Container, err error) {
	saveVal := "false"
	if safe {
		saveVal = "true"
	}
	return makeJSONCall(config.URL()+"/v1/snapshots/"+name, HTTPPut, "instance_id="+instance_id+"&safe="+saveVal)
}

func SnapshotDestroy(id string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/snapshots/"+id, HTTPDelete, "")
}

// Utility functions ---------------------------------------------------------------------------------------------------

func SnapshotFind(search string) string {
	ret := ""
	snapshots, err := SnapshotsList()
	if err != nil {
		fmt.Println("DEBUG: Returning early because err is", err)
		return ret
	}
	items, _ := snapshots.Children()
	for _, child := range items {
		name := child.S("name").Data().(string)
		if name == search {
			if ret != "" {
				return ""
			} else {
				ret = name
			}
		} else if strings.Contains(name, search) {
			if ret != "" {
				return ""
			} else {
				ret = name
			}
		}
	}
	return ret
}
