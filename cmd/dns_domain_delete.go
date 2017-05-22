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

var dnsDomainRemoveName string
var dnsDomainRemoveID string

var dnsDomainRemoveCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a domain",
	Aliases: []string{"destroy", "delete", "remove"},
	Long:    `Removed the domain with the name or ID specified`,
	Example: "civo domain remove --name=example.com",
	Run: func(cmd *cobra.Command, args []string) {
		if dnsDomainRemoveID == "" && dnsDomainRemoveName == "" {
			fmt.Println("You need to specify an ID or domain name")
			os.Exit(-1)
		}

		if dnsDomainRemoveID == "" {
			dnsDomainRemoveID = api.DnsDomainFind(dnsDomainRemoveName)
			if dnsDomainRemoveID == "" {
				fmt.Println("No domain matched with that name")
				os.Exit(-1)
			}
		}

		_, err := api.DnsDomainDestroy(dnsDomainRemoveID)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Println("Destroying domain", dnsDomainRemoveID)
	},
}

func init() {
	dnsDomainCmd.AddCommand(dnsDomainRemoveCmd)
	dnsDomainRemoveCmd.Flags().StringVarP(&dnsDomainRemoveID, "id", "i", "", "The domain name's ID")
	dnsDomainRemoveCmd.Flags().StringVarP(&dnsDomainRemoveName, "name", "n", "", "The domain name")

}
