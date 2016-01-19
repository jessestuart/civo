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

var tokenRemoveName string

// tokenRemoveCmd represents the accounts command
var tokenRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"delete", "forget"},
	Short:   "Remove a token",
	Long:    `To keep your list of tokens tidy, you can remove ones no longer necessary`,
	Run: func(cmd *cobra.Command, args []string) {
		if tokenRemoveName == "" {
			fmt.Println("You need to specify a name with --name in order to save the token")
			os.Exit(-3)
		}

		config.TokenRemove(tokenRemoveName)
		fmt.Printf("Removed token %s\n", tokenRemoveName)
	},
}

func init() {
	tokenCmd.AddCommand(tokenRemoveCmd)
	tokenRemoveCmd.Flags().StringVarP(&tokenRemoveName, "name", "n", "", "The name of the token to remove")
}
