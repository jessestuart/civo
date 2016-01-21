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

var ipConnectPublicIP string
var ipConnectPrivateIP string
var ipConnectInstance string

var ipConnectCommand = &cobra.Command{
	Use:     "connect",
	Aliases: []string{"attach", "link", "move"},
	Short:   "Connect/move a public IP",
	Example: "civo ip connect --public-ip 8.8.8.8 --private-ip 10.0.0.3 [name or ID]",
	Long:    `Connect an existing public IP to a private IP address on the specified instance (moving it from another instance)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need to specify a name or a partial/whole ID")
			os.Exit(-1)
		}

		if ipConnectPublicIP == "" {
			fmt.Println("You need to specify a public IP with --public-ip in order to reconnect that IP")
			os.Exit(-3)
		}

		if ipConnectPrivateIP == "" {
			fmt.Println("You need to specify a private IP with --private-ip in order to reconnect that IP")
			os.Exit(-3)
		}

		search := args[0]
		id := api.InstanceFind(search)
		if id == "" {
			fmt.Println("Couldn't find a single instance based on that name or partial/whole ID, it must match exactly one instance")
			os.Exit(-1)
		}

		_, err := api.IPConnect(id, ipConnectPublicIP, ipConnectPrivateIP)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}

		parts := strings.Split(id, "-")
		shortID := parts[0]
		fmt.Printf("Your request to connect %s to instance ID `%s` address %s was received and is being processed\n", ipConnectPublicIP, shortID, ipConnectPrivateIP)
	},
}

func init() {
	ipCommand.AddCommand(ipConnectCommand)
	ipConnectCommand.Flags().StringVarP(&ipConnectPublicIP, "public-ip", "u", "", "The public IP address to route to the specified instance")
	ipConnectCommand.Flags().StringVarP(&ipConnectPrivateIP, "private-ip", "r", "", "The private IP address on the instance to receive the traffic")
}
