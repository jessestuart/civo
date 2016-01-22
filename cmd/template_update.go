// Copyright © 2016 Absolute DevOps Ltd <info@absolutedevops.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/absolutedevops/civo/api"
	"github.com/absolutedevops/civo/config"
	"github.com/spf13/cobra"
)

var templateUpdateID string
var templateUpdateName string
var templateUpdateImageId string
var templateUpdateDescription string
var templateUpdateShortDescription string
var templateUpdateCloudInitFile string

var templateUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"save"},
	Short:   "Save a template",
	Example: "civo templates update --id ubuntu-70.04 --image-id b834b08c-bec6-11e5-b756-5cf9389be614",
	Long:    `Update a template with new parameters for people to use to build images from`,
	Run: func(cmd *cobra.Command, args []string) {
		if templateUpdateID == "" {
			fmt.Println("You need to specify a name with --id in order to update a template")
			os.Exit(-3)
		}

		params := api.TemplateParams{}
		res, err := api.TemplateFind(templateUpdateID)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		params.ID = res.Path("template.id").Data().(string)
		params.Name = res.Path("template.name").Data().(string)
		params.ImageID = res.Path("template.image_id").Data().(string)
		params.ShortDescription = res.Path("template.short_description").Data().(string)
		params.Description = res.Path("template.description").Data().(string)
		params.CloudConfig = res.Path("template.cloud_config").Data().(string)

		if templateUpdateName != "" {
			params.Name = templateUpdateName
		}
		if templateUpdateImageId != "" {
			params.ImageID = templateUpdateImageId
		}
		if templateUpdateShortDescription != "" {
			params.ShortDescription = templateUpdateShortDescription
		}
		if templateUpdateDescription != "" {
			params.Description = templateUpdateDescription
		}

		if templateUpdateCloudInitFile != "" {
			filename := strings.Replace(templateUpdateCloudInitFile, "~", "$HOME", -1)
			filename = os.ExpandEnv(filename)
			buf, err := ioutil.ReadFile(filename)
			if err == nil {
				params.CloudConfig = string(buf)
			}
		}

		_, err = api.TemplateUpdate(params)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		fmt.Printf("Updated template called `%s`\n", params.Name)
	},
}

func init() {
	if config.Admin() {
		templateCmd.AddCommand(templateUpdateCmd)
		templateUpdateCmd.Flags().StringVarP(&templateUpdateID, "id", "", "", "ID of the template; lowercase, hyphen separated")
		templateUpdateCmd.Flags().StringVarP(&templateUpdateName, "name", "n", "", "A nice name to be used for the template")
		templateUpdateCmd.Flags().StringVarP(&templateUpdateImageId, "image-id", "i", "", "The glance ID of the base filesystem image")
		templateUpdateCmd.Flags().StringVarP(&templateUpdateDescription, "description", "d", "", "A full/long multiline description")
		templateUpdateCmd.Flags().StringVarP(&templateUpdateShortDescription, "short-description", "s", "", "A one line short summary of the template")
		templateUpdateCmd.Flags().StringVarP(&templateUpdateCloudInitFile, "cloud-init-file", "c", "", "The filename of a file to be used as user-data/cloud-init")
	}
}
