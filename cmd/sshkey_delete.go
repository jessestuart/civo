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

var sshKeyDeleteName string

var sshKeyDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"new", "store"},
	Short:   "Upload an SSH key",
	Example: "civo sshkey upload --name default --public-key ~/.ssh/id_rsa.pub",
	Long:    `Upload an SSH public key, then this can be chosen when creating an instance`,
	Run: func(cmd *cobra.Command, args []string) {
		if sshKeyDeleteName == "" {
			fmt.Println("You need to specify a name with --name in order to remove that key")
			os.Exit(-3)
		}

		res, err := api.SshKeyDelete(sshKeyDeleteName)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Println(res)
		fmt.Printf("Deleted SSH key called `%s`\n", sshKeyDeleteName)
	},
}

func init() {
	sshKeyCommand.AddCommand(sshKeyDeleteCmd)
	sshKeyDeleteCmd.Flags().StringVarP(&sshKeyDeleteName, "name", "n", "", "The name of the key to be deleted")
}
