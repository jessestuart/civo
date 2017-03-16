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

	"github.com/absolutedevops/civo/config"
	"github.com/spf13/cobra"
)

var urlValue string

// urlCmd represents the accounts command
var urlCmd = &cobra.Command{
	Use:     "url",
	Aliases: []string{"api", "server"},
	Short:   "Change your current API URL",
	Long:    `Choose which URL the command line client should use in future connections`,
	Run: func(cmd *cobra.Command, args []string) {
		if urlValue == "" && len(args) > 0 {
			urlValue = args[0]
		}
		if urlValue == "" {
			fmt.Println("Resetting to the default URL...")
			urlValue = "https://api.civo.com"
		}

		config.APIKeySetURL(urlValue)
		fmt.Printf("Current URL is now %s\n", urlValue)
	},
}

func init() {
	if config.Admin() {
		RootCmd.AddCommand(urlCmd)
		urlCmd.Flags().StringVarP(&urlValue, "url", "u", "", "The URL of the API server to connect to")
	}
}
