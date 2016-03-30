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
	"github.com/spf13/cobra"
)

var ipDeleteIPAddress string
var ipDeleteInstance string

var ipDeleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "destroy"},
	Short:   "Remove an IP address",
	Example: "civo ip delete --ip-address 10.0.0.2 [name or ID]",
	Long:    `Remove a private IP address, or a private/public IP address pair (if specifying a public IP) from a specified instance`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need to specify a name or a partial/whole ID")
			os.Exit(-1)
		}

		search := args[0]
		id := api.InstanceFind(search)
		if id == "" {
			fmt.Println("Couldn't find a single instance based on that name or partial/whole ID, it must match exactly one instance")
			os.Exit(-1)
		}

		_, err := api.IPDelete(id, ipDeleteIPAddress)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}

		parts := strings.Split(id, "-")
		shortID := parts[0]
		fmt.Printf("Your request to remove %s from instance ID `%s` was received and is being processed\n", ipDeleteIPAddress, shortID)
	},
}

func init() {
	ipCommand.AddCommand(ipDeleteCommand)
	ipDeleteCommand.Flags().StringVarP(&ipDeleteIPAddress, "ip-address", "i", "", "The IP address to delete (can be private or public, but if private will also delete any attached public one)")
}
