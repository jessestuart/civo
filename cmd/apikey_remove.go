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

var apikeyRemoveName string

// apikeyRemoveCmd represents the accounts command
var apikeyRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"delete", "forget"},
	Short:   "Remove an API Key",
	Long:    `To keep your list of API Keys tidy, you can remove ones no longer necessary`,
	Run: func(cmd *cobra.Command, args []string) {
		if apikeyRemoveName == "" {
			fmt.Println("You need to specify a name with --name in order to save the apikey")
			os.Exit(-3)
		}

		config.APIKeyRemove(apikeyRemoveName)
		fmt.Printf("Removed API Key %s\n", apikeyRemoveName)
	},
}

func init() {
	apikeyCmd.AddCommand(apikeyRemoveCmd)
	apikeyRemoveCmd.Flags().StringVarP(&apikeyRemoveName, "name", "n", "", "The name of the API Key to remove")
}
