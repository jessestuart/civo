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
	"strings"

	"github.com/absolutedevops/civo/api"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var instanceFullIDs bool
var instanceTagsSearch string

// instanceCmd represents the accounts command
var instanceCmd = &cobra.Command{
	Use:     "instance",
	Aliases: []string{"instances"},
	Short:   "List all instances",
	Long:    `List the instances for the current account`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.InstancesList(instanceTagsSearch)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"ID", "Name", "Size", "Template", "IP Addresses", "Status", "User", "Password", "Firewall", "Tags"})
		items, _ := result.S("items").Children()
		for _, child := range items {
			var ipAddresses string
			privateIP := child.S("private_ip").Data().(string)
			publicIP := child.S("public_ip").Data().(string)
			if publicIP != "" {
				ipAddresses = privateIP + " => " + publicIP
			} else {
				ipAddresses = privateIP
			}

			var id string
			if instanceFullIDs {
				id = child.S("id").Data().(string)
			} else {
				parts := strings.Split(child.S("id").Data().(string), "-")
				id = parts[0]
			}

			tags := make([]string, 0)
			rawTags, _ := child.S("tags").Children()
			for _, rawTag := range rawTags {
				tags = append(tags, rawTag.Data().(string))
			}

			table.Append([]string{
				id,
				child.S("hostname").Data().(string),
				child.S("size").Data().(string),
				child.S("template").Data().(string),
				ipAddresses,
				child.S("status").Data().(string),
				child.S("initial_user").Data().(string),
				child.S("initial_password").Data().(string),
				child.S("firewall_id").Data().(string),
				strings.Join(tags, ", "),
			})
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(instanceCmd)
	instanceCmd.Flags().StringVarP(&instanceTagsSearch, "tags", "t", "", "Only return instances with these tags (AND not OR)")
	instanceCmd.Flags().BoolVarP(&instanceFullIDs, "full-ids", "f", false, "Return full IDs for instances")
}
