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
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var instanceFullIDs bool

// instanceCmd represents the accounts command
var instanceCmd = &cobra.Command{
	Use:     "instance",
	Aliases: []string{"instances"},
	Short:   "List all instances",
	Long:    `List the instances for the current account`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.InstancesList()
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"ID", "Name", "Size", "Template", "IP Addresses", "Status", "User", "Password"})
		items, _ := result.S("items").Children()
		for _, child := range items {
			ips, _ := child.S("ip_addresses").Children()
			ipAddresses := ""
			for _, ip := range ips {
				privateIP := ip.S("private_ip").Data().(string)
				publicIP := ip.S("public_ip").Data().(string)
				if ipAddresses != "" {
					ipAddresses = ipAddresses + ", "
				}
				if publicIP != "" {
					ipAddresses = ipAddresses + privateIP + "=>" + publicIP
				} else {
					ipAddresses = ipAddresses + privateIP
				}
			}

			var id string
			if instanceFullIDs {
				id = child.S("id").Data().(string)
			} else {
				parts := strings.Split(child.S("id").Data().(string), "-")
				id = parts[0]
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
			})
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(instanceCmd)
	instanceCmd.Flags().BoolVarP(&instanceFullIDs, "full-ids", "", false, "Return full IDs for instances")
}
