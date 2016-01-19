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
	"os"

	"github.com/absolutedevops/civo/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// tokenCmd represents the accounts command
var tokenCmd = &cobra.Command{
	Use:     "token",
	Aliases: []string{"tokens"},
	Short:   "List all tokens",
	Long:    `List the tokens you've saved for accessing the API`,
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"Name", "API Key", "Current"})
		for name, apiKey := range config.Tokens() {
			current := ""
			if name == config.TokenCurrent() {
				current = "<====="
			}
			table.Append([]string{
				name,
				apiKey,
				current,
			})
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(tokenCmd)
}
