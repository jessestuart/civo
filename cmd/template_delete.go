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
	"github.com/spf13/cobra"
)

var templateDestroyCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a template",
	Aliases: []string{"destroy", "delete", "remove"},
	Long:    `Given an ID that matches an account specific template, remove that template`,
	Example: "civo template remove [ID]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need to specify an ID")
			os.Exit(-1)
		}

		id := args[0]
		_, err := api.TemplateDestroy(id)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		fmt.Println("Destroying template ", id)
	},
}

func init() {
	templateCmd.AddCommand(templateDestroyCmd)
}
