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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var sshKeyDeleteName string
var sshKeyDeleteID string

var sshKeyDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"new", "store"},
	Short:   "Upload an SSH key",
	Example: "civo sshkey delete --id f0203e48",
	Long:    `Delete an SSH public key`,
	Run: func(cmd *cobra.Command, args []string) {
		errorColor := color.New(color.FgRed, color.Bold).SprintFunc()

		if sshKeyDeleteName == "" && sshKeyDeleteID == "" {
			fmt.Println("You need to specify an ID with --id or a name with --name in order to remove that key")
			os.Exit(-3)
		}

		keys, err := api.SshKeysList()
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}

		var foundID string
		var foundName string
		children, _ := keys.Children()
		for _, child := range children {
			if sshKeyDeleteName != "" && child.S("name").Data().(string) == sshKeyDeleteName {
				foundID = child.S("id").Data().(string)
				foundName = child.S("name").Data().(string)
			} else if sshKeyDeleteID != "" && strings.HasPrefix(child.S("id").Data().(string), sshKeyDeleteID) {
				foundID = child.S("id").Data().(string)
				foundName = child.S("name").Data().(string)
			}
		}

		if foundID != "" {
			res, err := api.SshKeyDelete(foundID)
			if err != nil {
				fmt.Println(errorColor("An error occured:"), err.Error())
				return
			}
			if res.S("result").Data() != nil && res.S("result").Data().(string) == "success" {
				fmt.Printf("Deleted SSH key called `%s`\n", foundName)
			} else {
				fmt.Println(errorColor("An error occured:"), res)
			}
		} else {
			fmt.Printf("Couldn't find that SSH key\n")
		}
	},
}

func init() {
	sshKeyCommand.AddCommand(sshKeyDeleteCmd)
	sshKeyDeleteCmd.Flags().StringVarP(&sshKeyDeleteName, "name", "n", "", "The name of the key to be deleted")
	sshKeyDeleteCmd.Flags().StringVarP(&sshKeyDeleteID, "id", "i", "", "The ID of the key to be deleted")
}
