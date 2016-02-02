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
	"github.com/spf13/cobra"
)

var snapshotDestroyCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a snapshot",
	Aliases: []string{"destroy", "delete", "remove"},
	Long:    `Given a name or partial/whole ID that matches one snapshot, remove that snapshot`,
	Example: "civo snapshot remove [name or ID]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need to specify a name or a partial/whole ID")
			os.Exit(-1)
		}

		search := args[0]
		id := api.SnapshotFind(search)
		if id == "" {
			fmt.Printf("Unable to find a snapshot matching '%s'", search)
			return
		}
		_, err := api.SnapshotDestroy(id)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		fmt.Println("Destroying snapshot ", search)
	},
}

func init() {
	snapshotCmd.AddCommand(snapshotDestroyCmd)
}
