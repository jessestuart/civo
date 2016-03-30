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

var instanceUpgradeSize string

var instanceUpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"resize"},
	Short:   "Upgrade an instance",
	Example: "civo instance upgrade --size g1.medium [name or ID]",
	Long:    `Upgrade the CPU, RAM and SSD disk for an instance with the specifed name or partial/full ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need to specify a name or a partial/whole ID")
			os.Exit(-1)
		}
		if instanceUpgradeSize == "" {
			fmt.Println("You need to specify a size with --size to specify the new size")
			os.Exit(-3)
		}

		search := args[0]
		id := api.InstanceFind(search)
		if id == "" {
			fmt.Println("Couldn't find a single instance based on that name or partial/whole ID, it must match exactly one instance")
			os.Exit(-1)
		}

		_, err := api.InstanceUpgrade(id, instanceUpgradeSize)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Resizing instance with ID %s to %s\n", id, instanceUpgradeSize)
	},
}

func init() {
	instanceCmd.AddCommand(instanceUpgradeCmd)
	instanceUpgradeCmd.Flags().StringVarP(&instanceUpgradeSize, "size", "s", "", "Upgrade the instance, using a size from `civo sizes`")
}
