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

var dnsRecordDeleteID string
var dnsRecordDeleteRecordID string

var dnsRecordDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove", "destroy"},
	Short:   "Delete a DNS record",
	Example: "civo domain delete --id={uuid} --record-id={uuid}",
	Long:    `Delete a DNS record for the specified domain with the given ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if dnsRecordDeleteID == "" || dnsRecordDeleteRecordID == "" {
			fmt.Println("You need to specify an ID for both the name and record")
			os.Exit(-1)
		}

		_, err := api.DnsRecordDelete(dnsRecordDeleteID, dnsRecordDeleteRecordID)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Removed DNS record '%s' from domain '%s'\n", dnsRecordDeleteRecordID, dnsRecordDeleteID)
	},
}

func init() {
	dnsRecordsCmd.AddCommand(dnsRecordDeleteCmd)
	dnsRecordDeleteCmd.Flags().StringVarP(&dnsRecordDeleteID, "id", "i", "", "ID of the firewall")
	dnsRecordDeleteCmd.Flags().StringVarP(&dnsRecordDeleteRecordID, "record-id", "r", "", "Which record ID to delete")
}
