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

var templateDestroyID string

var templateDestroyCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a template",
	Aliases: []string{"destroy", "delete", "remove"},
	Long:    `Given an ID that matches an account specific template, remove that template`,
	Example: "civo template remove [ID]",
	Run: func(cmd *cobra.Command, args []string) {
		if templateDestroyID == "" {
			fmt.Println("Couldn't remove a template without an ID")
			os.Exit(-1)
		}

		_, err := api.TemplateDestroy(templateDestroyID)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Println("Destroying template ", templateDestroyID)
	},
}

func init() {
	templateCmd.AddCommand(templateDestroyCmd)
	templateDestroyCmd.Flags().StringVarP(&templateDestroyID, "id", "i", "", "The network ID to delete")
}
