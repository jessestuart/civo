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
	"github.com/spf13/cobra"
)

var instanceFirewallInstanceID string
var instanceFirewallID string

var instanceFirewallCmd = &cobra.Command{
	Use:     "firewall",
	Aliases: []string{"revert"},
	Short:   "Firewall an instance",
	Example: "civo instance firewall --id {uuid} -f {uuid}",
	Long:    `Firewall an instance with the specifed firewall ID`,
	Run: func(cmd *cobra.Command, args []string) {
		instanceFirewallInstanceID := api.InstanceFind(instanceFirewallInstanceID)
		if instanceFirewallInstanceID == "" {
			fmt.Println("Couldn't find a single instance based on that name or partial/whole ID, it must match exactly one instance")
			os.Exit(-1)
		}

		_, err := api.InstanceFirewall(instanceFirewallInstanceID, instanceFirewallID)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Set firewall for instance %s to %s\n", instanceFirewallInstanceID, instanceFirewallID)
	},
}

func init() {
	instanceCmd.AddCommand(instanceFirewallCmd)
	instanceFirewallCmd.Flags().StringVarP(&instanceFirewallInstanceID, "id", "i", "", "The instance ID to use")
	instanceFirewallCmd.Flags().StringVarP(&instanceFirewallID, "firewall", "f", "", "The firewall to use")
}
