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

var firewallCreateName string
var firewallCreateInstanceID string
var firewallCreateSafe bool

var firewallCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add"},
	Short:   "Create a new firewall",
	Example: "civo firewall create --name restrictive",
	Long:    `Create a new firewall with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		if firewallCreateName == "" && len(args) > 0 {
			firewallCreateName = args[0]
		}

		_, err := api.FirewallCreate(firewallCreateName)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Creating a firewall with name '%s'\n", firewallCreateName)
	},
}

func init() {
	firewallCmd.AddCommand(firewallCreateCmd)
	firewallCreateCmd.Flags().StringVarP(&firewallCreateName, "name", "n", "", "Name of the firewall; lowercase, hyphen separated. If you don't specify one, a UUID followed by the instance_id will be used.")
}
