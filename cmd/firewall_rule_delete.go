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
	"github.com/spf13/cobra"
)

var firewallRuleDeleteName string
var firewallRuleDeleteRuleID string

var firewallRuleDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "destroy"},
	Short:   "Delete a firewall rule",
	Example: "civo firewall rule delete --name restrictive --id 1",
	Long:    `Delete a new firewall rule for the firewall with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		if firewallRuleDeleteName == "" && len(args) > 0 {
			firewallRuleDeleteName = args[0]
		}

		_, err := api.FirewallRuleDelete(firewallRuleDeleteName, firewallRuleDeleteRuleID)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		fmt.Printf("Removed firewall rule '%s' from firewall '%s'\n", firewallRuleDeleteRuleID, firewallRuleDeleteName)
	},
}

func init() {
	firewallRuleCmd.AddCommand(firewallRuleDeleteCmd)
	firewallRuleDeleteCmd.Flags().StringVarP(&firewallRuleDeleteName, "name", "n", "", "Name of the firewall; lowercase, hyphen separated. If you don't specify one, a UUID followed by the instance_id will be used.")
	firewallRuleDeleteCmd.Flags().StringVarP(&firewallRuleDeleteRuleID, "id", "i", "", "Which rule ID to delete")
}
