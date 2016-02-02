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
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// snapshotCmd represents the accounts command
var snapshotCmd = &cobra.Command{
	Use:     "snapshot",
	Aliases: []string{"snapshots"},
	Short:   "List all snapshots",
	Long:    `List the snapshots for the current account`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.SnapshotsList()
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader([]string{"Name", "Instance", "Requested At", "Completed At", "Status", "Progress"})
		items, _ := result.S("items").Children()
		for _, child := range items {
			parts := strings.Split(child.S("instance_id").Data().(string), "-")
			instanceId := parts[0]

			table.Append([]string{
				child.S("name").Data().(string),
				instanceId,
				child.S("requested_at").Data().(string),
				child.S("completed_at").Data().(string),
				child.S("status").Data().(string),
				fmt.Sprintf("%.0f", child.S("progress").Data().(float64)) + "%",
			})
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(snapshotCmd)
}
