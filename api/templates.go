package api

import (
	"github.com/absolutedevops/civo/config"
	"github.com/google/go-querystring/query"
	"github.com/jeffail/gabs"
)

type TemplateParams struct {
	ID               string `url:"id"`
	Name             string `url:"name"`
	ImageID          string `url:"image_id"`
	ShortDescription string `url:"short_description"`
	Description      string `url:"description"`
	CloudConfig      string `url:"cloud_config"`
}

func TemplatesList() (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/templates", HTTPGet, "")
}

func TemplateCreate(params TemplateParams) (json *gabs.Container, err error) {
	v, _ := query.Values(params)
	return makeJSONCall(config.URL()+"/v1/templates", HTTPPost, v.Encode())
}

func TemplateUpdate(params TemplateParams) (json *gabs.Container, err error) {
	v, _ := query.Values(params)
	return makeJSONCall(config.URL()+"/v1/templates/"+params.ID, HTTPPut, v.Encode())
}

func TemplateFind(id string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/templates/"+id, HTTPGet, "")
}

func TemplateDestroy(id string) (json *gabs.Container, err error) {
	return makeJSONCall(config.URL()+"/v1/templates/"+id, HTTPDelete, "")
}
