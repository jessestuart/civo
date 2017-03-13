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

	"github.com/absolutedevops/civo/api"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var firewallRuleDeleteID string
var firewallRuleDeleteRuleID string

var firewallRuleDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "destroy"},
	Short:   "Delete a firewall rule",
	Example: "civo firewall rule delete --id {uuid} --rule-id {uuid}",
	Long:    `Delete a firewall rule for the firewall with the given ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if firewallRuleDeleteID == "" && len(args) > 0 {
			firewallRuleDeleteID = args[0]
		}

		_, err := api.FirewallRuleDelete(firewallRuleDeleteID, firewallRuleDeleteRuleID)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Removed firewall rule '%s' from firewall '%s'\n", firewallRuleDeleteRuleID, firewallRuleDeleteID)
	},
}

func init() {
	firewallRuleCmd.AddCommand(firewallRuleDeleteCmd)
	firewallRuleDeleteCmd.Flags().StringVarP(&firewallRuleDeleteID, "id", "i", "", "ID of the firewall")
	firewallRuleDeleteCmd.Flags().StringVarP(&firewallRuleDeleteRuleID, "rule-id", "r", "", "Which rule ID to delete")
}
