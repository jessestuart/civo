// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/absolutedevops/civo/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// accountCmd represents the accounts command
var accountCmd = &cobra.Command{
	Use:     "account",
	Aliases: []string{"accounts"},
	Short:   "List current accounts (ADMIN ONLY)",
	Long:    `List the account name and the API keys for all accounts in the system`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.AccountsList()
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Account Name", "API Key"})
		items, _ := result.S("items").Children()
		for _, child := range items {
			table.Append([]string{
				child.S("username").Data().(string),
				child.S("api_key").Data().(string),
			})
		}
		table.Render()
	},
}

func init() {
	if config.Admin() {
		RootCmd.AddCommand(accountCmd)
	}
}
