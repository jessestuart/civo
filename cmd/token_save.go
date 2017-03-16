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

	"github.com/absolutedevops/civo/config"
	"github.com/spf13/cobra"
)

var tokenSaveName string
var tokenSaveKey string

// tokenSaveCmd represents the accounts command
var tokenSaveCmd = &cobra.Command{
	Use:     "save",
	Aliases: []string{"register", "store"},
	Short:   "Save a token",
	Long:    `When you've been given an API key for use against a Civo cloud, you can save the token here to be able to reuse it within this client`,
	Run: func(cmd *cobra.Command, args []string) {
		if tokenSaveName == "" {
			fmt.Println("You need to specify a name with --name in order to save the token")
			os.Exit(-3)
		}
		if tokenSaveKey == "" {
			fmt.Println("You need to specify a key with --key in order to save the token")
			os.Exit(-3)
		}

		config.TokenSave(tokenSaveName, tokenSaveKey)
		if config.TokenCurrent() == "" {
			config.TokenSetCurrent(tokenSaveName)
			fmt.Printf("Saved token %s and set it as the default token because it's your first token\n", tokenSaveName)
		} else {
			fmt.Printf("Saved token %s, to use it as the default use : civo tokens default --name %s\n", tokenSaveName, tokenSaveName)
		}
	},
}

func init() {
	tokenCmd.AddCommand(tokenSaveCmd)
	tokenSaveCmd.Flags().StringVarP(&tokenSaveName, "name", "n", "", "The name to use for this token (can be an abbreviation)")
	tokenSaveCmd.Flags().StringVarP(&tokenSaveKey, "key", "k", "", "The API key supplied for use against a Civo cloud")
}
