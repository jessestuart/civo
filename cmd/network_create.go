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

var networkCreateName string
var networkCreateLabel string
var networkCreateRegion string

var networkCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add"},
	Short:   "Create a new network",
	Example: "civo network create --name development",
	Long:    `Create a new network with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		if networkCreateName == "" && len(args) > 0 {
			networkCreateName = args[0]
		}

		if networkCreateLabel == "" {
			networkCreateLabel = networkCreateName
		}

		_, err := api.NetworkCreate(networkCreateName, networkCreateLabel, networkCreateRegion)
		if err != nil {
			errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(errorColor("An error occured:"), err.Error())
			return
		}
		fmt.Printf("Created a network with name '%s'\n", networkCreateName)
	},
}

func init() {
	networkCmd.AddCommand(networkCreateCmd)
	networkCreateCmd.Flags().StringVarP(&networkCreateName, "name", "n", "", "Name of the network; lowercase, hyphen separated. If you don't specify one, a UUID will be used.")
	networkCreateCmd.Flags().StringVarP(&networkCreateLabel, "label", "l", "", "A nice name for the network; can contain spaces, etc.")
	networkCreateCmd.Flags().StringVarP(&networkCreateRegion, "region", "r", DefaultRegion, "The region from 'civo regions'")
}
