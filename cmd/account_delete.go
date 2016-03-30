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
	"github.com/absolutedevops/civo/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var accountDeleteName string

var accountDeleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete an account",
	Aliases: []string{"destroy", "remove"},
	Long:    `Given a name, delete its account and all instances`,
	Run: func(cmd *cobra.Command, args []string) {
		if accountDeleteName == "" {
			fmt.Println("You need to specify a name with --name in order to delete an account")
			os.Exit(-3)
		}

		_, err := api.AccountDelete(accountDeleteName)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
	},
}

func init() {
	if config.Admin() {
		accountCmd.AddCommand(accountDeleteCmd)
		accountDeleteCmd.Flags().StringVarP(&accountDeleteName, "name", "n", "", "Name of the account to delete")
	}
}
