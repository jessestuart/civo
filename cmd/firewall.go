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
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// firewallCmd represents the accounts command
var firewallCmd = &cobra.Command{
	Use:     "firewall",
	Aliases: []string{"firewalls"},
	Short:   "List all firewalls",
	Long:    `List the firewalls for the current account`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.FirewallsList()
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"Name", "Instances", "Rules"})
		items, _ := result.S("items").Children()
		for _, child := range items {
			table.Append([]string{
				child.S("name").Data().(string),
				fmt.Sprintf("%.0f", child.S("instances_count").Data().(float64)),
				fmt.Sprintf("%.0f", child.S("rules_count").Data().(float64)),
			})
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(firewallCmd)
}
