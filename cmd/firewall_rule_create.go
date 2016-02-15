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

var firewallRuleCreateName string
var firewallRuleCreateProtocol string
var firewallRuleCreateStartPort string
var firewallRuleCreateEndPort string
var firewallRuleCreateCIDR string
var firewallRuleCreateDirection string

var firewallRuleCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add"},
	Short:   "Create a new firewall rule",
	Example: "civo firewall rule create --name restrictive",
	Long:    `Create a new firewall rule for the firewall with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		if firewallRuleCreateName == "" && len(args) > 0 {
			firewallRuleCreateName = args[0]
		}

		params := api.FirewallRuleParams{
			Protocol:  firewallRuleCreateProtocol,
			StartPort: firewallRuleCreateStartPort,
			EndPort:   firewallRuleCreateEndPort,
			CIDR:      firewallRuleCreateCIDR,
			Direction: firewallRuleCreateDirection,
		}

		result, err := api.FirewallRuleCreate(firewallRuleCreateName, params)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		fmt.Printf("Creating a firewall rule to allow '%s' %s access to ports '%s/%s' on firewall '%s' with ID '%d'\n", firewallRuleCreateCIDR, firewallRuleCreateDirection, firewallRuleCreateStartPort, firewallRuleCreateEndPort, firewallRuleCreateName, fmt.Sprintf("%.0f", child.S("id").Data().(float64)))
	},
}

func init() {
	firewallRuleCmd.AddCommand(firewallRuleCreateCmd)
	firewallRuleCreateCmd.Flags().StringVarP(&firewallRuleCreateName, "name", "n", "", "Name of the firewall; lowercase, hyphen separated. If you don't specify one, a UUID followed by the instance_id will be used.")
	firewallRuleCreateCmd.Flags().StringVarP(&firewallRuleCreateProtocol, "protocol", "p", "tcp", "Which internet protocol to filter: tcp, udp or icmp")
	firewallRuleCreateCmd.Flags().StringVarP(&firewallRuleCreateStartPort, "start", "s", "", "The start of the port range to allow")
	firewallRuleCreateCmd.Flags().StringVarP(&firewallRuleCreateEndPort, "end", "e", "", "The end of the port range to allow (either a different number to start, the same number or empty for allowing a single port)")
	firewallRuleCreateCmd.Flags().StringVarP(&firewallRuleCreateCIDR, "cidr", "c", "0.0.0.0/0", "The IP address or CIDR to filter")
	firewallRuleCreateCmd.Flags().StringVarP(&firewallRuleCreateDirection, "direction", "d", "inbound", "Will this rule affect inbound or outbound traffic")
}
