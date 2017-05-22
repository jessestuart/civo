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

var dnsRecordCreateID string
var dnsRecordCreateName string
var dnsRecordCreateValue string
var dnsRecordCreateType string
var dnsRecordCreatePriority string
var dnsRecordCreateTTL string

var dnsRecordCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add"},
	Short:   "Create a new DNS record",
	Example: "civo dns record create --id={uuid} --name=www --type=a --value=1.2.3.4 ttl=600",
	Long:    `Create a new DNS record for the domain name with the given ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if dnsRecordCreateID == "" {
			fmt.Println("You need to specify an ID for the domain name")
			os.Exit(-1)
		}

		params := api.DnsRecordParams{
			Name:     dnsRecordCreateName,
			Value:    dnsRecordCreateValue,
			Type:     dnsRecordCreateType,
			Priority: dnsRecordCreatePriority,
			TTL:      dnsRecordCreateTTL,
		}

		result, err := api.DnsRecordCreate(dnsRecordCreateID, params)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Creating a DNS record to point '%s' to '%s' (type: %s)  with a TTL of %ss (%s)'\n", dnsRecordCreateName, dnsRecordCreateValue, dnsRecordCreateType, dnsRecordCreateTTL, result.S("id").Data().(string))
	},
}

func init() {
	dnsRecordsCmd.AddCommand(dnsRecordCreateCmd)
	dnsRecordCreateCmd.Flags().StringVarP(&dnsRecordCreateID, "id", "i", "", "ID of the domain")
	dnsRecordCreateCmd.Flags().StringVarP(&dnsRecordCreateName, "name", "n", "", "Name of the record, e.g. 'www'")
	dnsRecordCreateCmd.Flags().StringVarP(&dnsRecordCreateValue, "value", "v", "", "The value for the record, e.g. '1.2.3.4'")
	dnsRecordCreateCmd.Flags().StringVarP(&dnsRecordCreateType, "type", "t", "", "The type of DNS record (choice: a, cname, mx, txt)")
	dnsRecordCreateCmd.Flags().StringVarP(&dnsRecordCreatePriority, "priority", "p", "", "The priority for this value, if you're creating an 'mx' record")
	dnsRecordCreateCmd.Flags().StringVarP(&dnsRecordCreateTTL, "ttl", "l", "3600", "How long should DNS servers cache this value for")
}
