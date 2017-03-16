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

var apikeySaveName string
var apikeySaveKey string

// apikeySaveCmd represents the accounts command
var apikeySaveCmd = &cobra.Command{
	Use:     "save",
	Aliases: []string{"register", "store"},
	Short:   "Save an API Key",
	Long:    `When you've been given an API key for use against a Civo cloud API, you can save the API Key here to be able to use it within this client`,
	Run: func(cmd *cobra.Command, args []string) {
		if apikeySaveName == "" {
			fmt.Println("You need to specify a name with --name in order to save the apikey")
			os.Exit(-3)
		}
		if apikeySaveKey == "" {
			fmt.Println("You need to specify a key with --key in order to save the apikey")
			os.Exit(-3)
		}

		config.APIKeySave(apikeySaveName, apikeySaveKey)
		if config.APIKeyCurrent() == "" {
			config.APIKeySetCurrent(apikeySaveName)
			fmt.Printf("Saved API Key %s and set it as the default API Key because it's your first API Key\n", apikeySaveName)
		} else {
			fmt.Printf("Saved API Key %s, to use it as the default use : civo apikeys default --name %s\n", apikeySaveName, apikeySaveName)
		}
	},
}

func init() {
	apikeyCmd.AddCommand(apikeySaveCmd)
	apikeySaveCmd.Flags().StringVarP(&apikeySaveName, "name", "n", "", "The name to use for this API Key (can be an abbreviation)")
	apikeySaveCmd.Flags().StringVarP(&apikeySaveKey, "key", "k", "", "The API key supplied for use against a Civo cloud")
}
