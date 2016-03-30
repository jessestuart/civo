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
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// templateCmd represents the accounts command
var templateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"templates"},
	Short:   "List all templates",
	Long:    `List the templates available for building instances from`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.TemplatesList()
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"ID", "Description", "Global"})
		items, _ := result.Children()
		for _, child := range items {
			global := "yes"
			if child.S("tenant").Data().(string) != "" {
				global = "no"
			}
			table.Append([]string{
				child.S("id").Data().(string),
				child.S("short_description").Data().(string),
				global,
			})
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(templateCmd)
}
