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
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var dnsRecordsDomainID string

// dnsRecordCmd represents the accounts command
var dnsRecordsCmd = &cobra.Command{
	Use:     "records",
	Aliases: []string{"record"},
	Short:   "List all DNS records",
	Example: "civo domain records --id={uuid}",
	Long:    `List the records for the specified DNS domain name`,
	Run: func(cmd *cobra.Command, args []string) {
		if dnsRecordsDomainID == "" {
			fmt.Println("You need to specify a domain ID")
			os.Exit(-1)
		}

		result, err := api.DnsRecords(dnsRecordsDomainID)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"ID", "Name", "Value", "Type", "Priority", "TTL"})
		items, _ := result.Children()
		for _, child := range items {
			priority := fmt.Sprintf("%.0f", (child.S("priority").Data().(float64)))
			recordType := (child.S("type").Data().(string))
			if recordType != "mx" {
				priority = ""
			}
			name := (child.S("name").Data().(string))
			ttl := (child.S("ttl").Data().(float64))
			value := (child.S("value").Data().(string))

			table.Append([]string{
				child.S("id").Data().(string),
				name,
				value,
				strings.ToUpper(recordType),
				priority,
				fmt.Sprintf("%.0f", ttl),
			})
		}
		table.Render()
	},
}

func init() {
	dnsDomainCmd.AddCommand(dnsRecordsCmd)
	dnsRecordsCmd.Flags().StringVarP(&dnsRecordsDomainID, "id", "i", "", "ID of the domain name")
}
