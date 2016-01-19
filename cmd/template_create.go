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

var templateCreateID string
var templateCreateName string
var templateCreateImageId string
var templateCreateDescription string
var templateCreateShortDescription string
var templateCreateCloudInitFile string

var templateCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "build"},
	Short:   "Create a new template",
	Example: "civo templates create --name ubuntu-70.04 --image-id b834b08c-bec6-11e5-b756-5cf9389be614 --short-description 'An Ubuntu 70.04 base image'",
	Long:    `Create a new template for people to use to build images from`,
	Run: func(cmd *cobra.Command, args []string) {
		if templateCreateID == "" {
			fmt.Println("You need to specify a name with --name in order to create a template")
			os.Exit(-3)
		}
		if templateCreateName == "" {
			fmt.Println("You need to specify a name with --name in order to create a template")
			os.Exit(-3)
		}
		if templateCreateImageId == "" {
			fmt.Println("You need to specify a Glance image ID with --image-id in order to create a template")
			os.Exit(-3)
		}
		if templateCreateShortDescription == "" {
			fmt.Println("You need to specify a one-line description with --short-description in order to create a template")
			os.Exit(-3)
		}

		params := api.TemplateParams{
			ID:               templateCreateID,
			Name:             templateCreateName,
			ImageID:          templateCreateImageId,
			ShortDescription: templateCreateShortDescription,
			Description:      templateCreateDescription,
		}
		if templateCreateCloudInitFile != "" {
			filename := strings.Replace(templateCreateCloudInitFile, "~", "$HOME", -1)
			filename = os.ExpandEnv(filename)
			buf, err := ioutil.ReadFile(filename)
			if err == nil {
				params.CloudConfig = string(buf)
			}
		}

		_, err := api.TemplateCreate(params)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		fmt.Printf("Created template called `%s`\n", templateCreateName)
	},
}

func init() {
	if config.Admin() {
		templateCmd.AddCommand(templateCreateCmd)
		templateCreateCmd.Flags().StringVarP(&templateCreateID, "id", "", "", "ID of the template; lowercase, hyphen separated")
		templateCreateCmd.Flags().StringVarP(&templateCreateName, "name", "n", "", "A nice name to be used for the template")
		templateCreateCmd.Flags().StringVarP(&templateCreateImageId, "image-id", "i", "", "The glance ID of the base filesystem image")
		templateCreateCmd.Flags().StringVarP(&templateCreateDescription, "description", "d", "", "A full/long multiline description")
		templateCreateCmd.Flags().StringVarP(&templateCreateShortDescription, "short-description", "s", "", "A one line short summary of the template")
		templateCreateCmd.Flags().StringVarP(&templateCreateCloudInitFile, "cloud-init-file", "c", "", "The filename of a file to be used as user-data/cloud-init")
	}
}
