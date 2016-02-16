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
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// sizeCmd represents the accounts command
var sizeCmd = &cobra.Command{
	Use:     "size",
	Aliases: []string{"sizes"},
	Short:   "List sizes",
	Long:    `List the available sizes with their specifications for instance creation`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.SizesList()
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"Name", "Specification"})
		table.SetAutoWrapText(false)
		items, _ := result.Children()
		for _, child := range items {
			table.Append([]string{
				child.S("name").Data().(string),
				child.S("description").Data().(string),
			})
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(sizeCmd)
}
