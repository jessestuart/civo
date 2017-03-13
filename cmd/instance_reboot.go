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

var instanceRebootHard bool
var instanceRebootInstanceID string

var instanceRebootCmd = &cobra.Command{
	Use:     "reboot",
	Aliases: []string{"restart"},
	Short:   "Reboot an instance",
	Example: "civo instance reboot --id {uuid}",
	Long:    `Reboot an instance with the specifed name or partial/full ID`,
	Run: func(cmd *cobra.Command, args []string) {
		instanceRebootInstanceID := api.InstanceFind(instanceRebootInstanceID)
		if instanceRebootInstanceID == "" {
			fmt.Println("Couldn't find a single instance based on that name or partial/whole ID, it must match exactly one instance")
			os.Exit(-1)
		}

		_, err := api.InstanceReboot(instanceRebootInstanceID, instanceRebootHard)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Println("Rebooting instance with ID", instanceRebootInstanceID)
	},
}

func init() {
	instanceCmd.AddCommand(instanceRebootCmd)
	instanceRebootCmd.Flags().BoolVarP(&instanceRebootHard, "hard", "", false, "Perform a hard-reboot - literally kill the machine and restart it")
	instanceRebootCmd.Flags().StringVarP(&instanceRebootInstanceID, "id", "i", "", "The instance ID to reboot")

}
