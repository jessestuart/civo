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

var firewallRulesID string

// firewallRuleCmd represents the accounts command
var firewallRuleCmd = &cobra.Command{
	Use:     "rules",
	Aliases: []string{"rule"},
	Short:   "List all firewall rules",
	Example: "civo firewall rules --id {uuid}",
	Long:    `List the firewall rules for the specified firewall`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 && firewallRulesID == "" {
			fmt.Println("You need to specify a firewall ID")
			os.Exit(-1)
		}
		if firewallRulesID == "" {
			firewallRulesID = args[0]
		}

		result, err := api.FirewallRules(firewallRulesID)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"ID", "Protocol", "Port", "CIDR", "Direction"})
		items, _ := result.Children()
		for _, child := range items {
			startPort := (child.S("start_port").Data().(string))
			endPort := (child.S("end_port").Data().(string))
			ports := ""
			if startPort == endPort || endPort == "" {
				ports = startPort
			} else {
				ports = startPort + "-" + endPort
			}

			table.Append([]string{
				child.S("id").Data().(string),
				strings.ToUpper(child.S("protocol").Data().(string)),
				ports,
				child.S("cidr").Data().(string),
				child.S("direction").Data().(string),
			})
		}
		table.Render()
	},
}

func init() {
	firewallCmd.AddCommand(firewallRuleCmd)
	firewallRuleCmd.Flags().StringVarP(&firewallRulesID, "id", "i", "", "ID of the firewall")
}
