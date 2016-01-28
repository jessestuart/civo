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
	"os"

	"github.com/absolutedevops/civo/api"
	"github.com/absolutedevops/civo/config"
	"github.com/spf13/cobra"
)

var templateDetailsID string

var templateDetailsCmd = &cobra.Command{
	Use:     "details",
	Aliases: []string{"show", "info"},
	Short:   "Show full details about a template",
	Example: "civo templates details --name ubuntu-70.04",
	Long:    `Show the full details for a template`,
	Run: func(cmd *cobra.Command, args []string) {
		if templateDetailsID == "" {
			fmt.Println("You need to specify an ID with --id in order to show a template's full details")
			os.Exit(-3)
		}

		res, err := api.TemplateFind(templateDetailsID)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		// fmt.Println(res.String())
		fmt.Printf("%-20s: %s\n", "ID", res.Path("id").Data().(string))
		fmt.Printf("%-20s: %s\n", "Name", res.Path("name").Data().(string))
		fmt.Printf("%-20s: %s\n", "Image ID", res.Path("image_id").Data().(string))
		fmt.Printf("%-20s: %s\n", "Short Description", res.Path("short_description").Data().(string))
		fmt.Printf("%-20s: %s\n", "Description", res.Path("description").Data().(string))
		fmt.Println("")
		fmt.Println(">>>>> Cloud Config <<<<<")
		fmt.Println(res.Path("cloud_config").Data().(string))
	},
}

func init() {
	if config.Admin() {
		templateCmd.AddCommand(templateDetailsCmd)
		templateDetailsCmd.Flags().StringVarP(&templateDetailsID, "id", "", "", "ID of the template to show")
	}
}
