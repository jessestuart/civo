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

var instanceTagsForRetagging string
var instanceTagsInstanceID string

var instanceTagsCmd = &cobra.Command{
	Use:     "tags",
	Aliases: []string{"tag", "tagged", "retag"},
	Short:   "Re-tag an instance",
	Example: "civo instance tags --id {uuid} --tags \"[tag names, space separated]\"",
	Long:    `Re-tag an instance with the specifed name or partial/full ID using the list of space separated tags`,
	Run: func(cmd *cobra.Command, args []string) {
		instanceTagsInstanceID := api.InstanceFind(instanceTagsInstanceID)
		if instanceTagsInstanceID == "" {
			fmt.Println("Couldn't find a single instance based on that name or partial/whole ID, it must match exactly one instance")
			os.Exit(-1)
		}

		_, err := api.InstanceTags(instanceTagsInstanceID, instanceTagsForRetagging)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Set tags for instance %s to '%s'\n", instanceTagsInstanceID, instanceTagsForRetagging)
	},
}

func init() {
	instanceCmd.AddCommand(instanceTagsCmd)
	instanceTagsCmd.Flags().StringVarP(&instanceTagsForRetagging, "tags", "t", "", "The space separated list of tags")
	instanceTagsCmd.Flags().StringVarP(&instanceTagsInstanceID, "id", "i", "", "The instance ID to reboot")
}
