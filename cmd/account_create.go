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

var accountCreateName string

var accountCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "build", "register"},
	Short:   "Create a new account",
	Example: "create --name testuser",
	Long:    `Given a name, create an account with a new API key for it`,
	Run: func(cmd *cobra.Command, args []string) {
		if accountCreateName == "" {
			fmt.Println("You need to specify a name with --name in order to create an account")
			os.Exit(-3)
		}

		_, err := api.AccountCreate(accountCreateName)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
	},
}

func init() {
	accountCmd.AddCommand(accountCreateCmd)
	accountCreateCmd.Flags().StringVarP(&accountCreateName, "name", "n", "", "Name of the account; lowercase, hyphen separated")
}
