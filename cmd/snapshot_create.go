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

var snapshotCreateName string
var snapshotCreateInstanceID string
var snapshotCreateSafe bool

var snapshotCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "backup"},
	Short:   "Create a new snapshot",
	Example: "civo snapshot create --name my-backup --instance my-host.example.com",
	Long:    `Create a new snapshot of the specified instance with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		if snapshotCreateName == "" && len(args) > 0 {
			snapshotCreateName = args[0]
		}
		if snapshotCreateInstanceID == "" && snapshotCreateName != "" && len(args) > 0 {
			snapshotCreateInstanceID = args[0]
		} else if len(args) > 1 {
			snapshotCreateInstanceID = args[1]
		}

		instanceID := api.InstanceFind(snapshotCreateInstanceID)
		if instanceID == "" {
			fmt.Println("Couldn't find a single instance based on that name or partial/whole ID, it must match exactly one instance")
			os.Exit(-1)
		}

		_, err := api.SnapshotCreate(snapshotCreateName, instanceID, snapshotCreateSafe)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Creating a snapshot of `%s` with name '%s'\n", instanceID, snapshotCreateName)
	},
}

func init() {
	snapshotCmd.AddCommand(snapshotCreateCmd)
	snapshotCreateCmd.Flags().StringVarP(&snapshotCreateName, "name", "n", "", "Name of the snapshot; lowercase, hyphen separated. If you don't specify one, a UUID followed by the instance_id will be used.")
	snapshotCreateCmd.Flags().StringVarP(&snapshotCreateInstanceID, "instance", "i", "", "The ID or the hostname of the instance to snapshot")
	snapshotCreateCmd.Flags().BoolVarP(&snapshotCreateSafe, "safe", "s", false, "Whether to shutdown the instance beforehand, to ensure a safe snapshot")
}
