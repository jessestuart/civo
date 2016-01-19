// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
)

var accountResetName string

var accountResetCmd = &cobra.Command{
	Use:     "reset",
	Short:   "Reset the API key for an account",
	Example: "create --name testuser",
	Long:    `Given a name, reset the account by setting a new API key for it`,
	Run: func(cmd *cobra.Command, args []string) {
		if accountResetName == "" {
			fmt.Println("You need to specify a name with --name in order to reset an account")
			os.Exit(-3)
		}

		_, err := api.AccountReset(accountResetName)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
	},
}

func init() {
	if config.Admin() {
		accountCmd.AddCommand(accountResetCmd)
		accountResetCmd.Flags().StringVarP(&accountResetName, "name", "n", "", "Name of the account; lowercase, hyphen separated")
	}
}
