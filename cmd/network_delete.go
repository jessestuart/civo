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

var networkDestroyID string

var networkDestroyCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a network",
	Aliases: []string{"destroy", "delete", "remove"},
	Long:    `Given a name or partial/whole ID that matches one network, remove that network`,
	Example: "civo network remove --id {uuid}",
	Run: func(cmd *cobra.Command, args []string) {
		networkDestroyID := api.NetworkFind(networkDestroyID)
		if networkDestroyID == "" {
			fmt.Println("Couldn't find a single network based on that name or partial/whole ID, it must match exactly one network")
			os.Exit(-1)
		}

		_, err := api.NetworkDestroy(networkDestroyID)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Println("Destroying network with ID", networkDestroyID)
	},
}

func init() {
	networkCmd.AddCommand(networkDestroyCmd)
	networkDestroyCmd.Flags().StringVarP(&networkDestroyID, "id", "i", "", "The network ID to delete")
}
