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
	"github.com/spf13/cobra"
)

var ipCreatePublic bool
var ipCreateInstance string

var ipCreateCommand = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "allocate"},
	Short:   "Add a new IP address",
	Example: "civo ip add --public [name or ID]",
	Long:    `Add a new private IP address, or a private/public IP address pair to a specified instance`,
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

		result, err := api.IPCreate(id, ipCreatePublic)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}

		parts := strings.Split(id, "-")
		shortID := parts[0]
		status := result.S("result").Data().(string)
		if status == "pending" {
			fmt.Printf("Your request to create an IP address for instance ID `%s` was received and is being processed\n", shortID)
		} else {
			privateIP := result.S("private_ip").Data().(string)
			var publicIP string
			if result.S("public_ip") != nil {
				publicIP = result.S("public_ip").Data().(string)
			}
			if publicIP != "" {
				fmt.Printf("Your IP address has been created; private is %s and public is %s\n", privateIP, publicIP)
			} else {
				fmt.Printf("Your IP address has been created - %s\n", privateIP)
				if ipCreatePublic {
					fmt.Println("The public IP address is being procesed, please check - civo instances - in a few minutes")
				}
			}
		}
	},
}

func init() {
	ipCommand.AddCommand(ipCreateCommand)
	ipCreateCommand.Flags().BoolVarP(&ipCreatePublic, "public", "p", true, "Should this create a public IP address, along with the new private one or just a private IP address")
}
