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

	"github.com/absolutedevops/civo/api"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var dnsDomainCreateName string

var dnsDomainCreateCmd = &cobra.Command{
	Use:     "setup",
	Aliases: []string{"create", "new", "add"},
	Short:   "Setup a new domain",
	Example: "civo domain setup --name=restrictive",
	Long:    `Setup DNS hosting for a domain with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		if dnsDomainCreateName == "" && len(args) > 0 {
			dnsDomainCreateName = args[0]
		}

		_, err := api.DnsDomainCreate(dnsDomainCreateName)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Setting up DNS hosting for the domain '%s'\n", dnsDomainCreateName)
	},
}

func init() {
	dnsDomainCmd.AddCommand(dnsDomainCreateCmd)
	dnsDomainCreateCmd.Flags().StringVarP(&dnsDomainCreateName, "name", "n", "", "The domain name")
}
