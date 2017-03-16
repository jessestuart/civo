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

var apikeyCurrentName string

// apikeyCurrentCmd represents the accounts command
var apikeyCurrentCmd = &cobra.Command{
	Use:     "current",
	Aliases: []string{"default", "use", "choose", "select"},
	Short:   "Change your current API Key",
	Long:    `Choose which API Key the command line client should use in future connections`,
	Run: func(cmd *cobra.Command, args []string) {
		if apikeyCurrentName == "" && len(args) > 0 {
			apikeyCurrentName = args[0]
		}
		if apikeyCurrentName == "" {
			fmt.Println("You need to specify a name with --name in order to set the API Key as the current one")
			os.Exit(-3)
		}

		config.APIKeySetCurrent(apikeyCurrentName)
		fmt.Printf("Current API Key is now %s\n", apikeyCurrentName)
	},
}

func init() {
	apikeyCmd.AddCommand(apikeyCurrentCmd)
	apikeyCurrentCmd.Flags().StringVarP(&apikeyCurrentName, "name", "n", "", "The name to use for this apikey (can be an abbreviation)")
}
