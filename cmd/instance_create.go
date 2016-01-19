// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
)

var instanceCreateName string
var instanceCreateSize string
var instanceCreateRegion string
var instanceCreateSSHKey string
var instanceCreatePublicIP bool
var instanceCreateTemplate string
var instanceCreateInitialUser string

var instanceCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "build"},
	Short:   "Create a new instance",
	Example: "civo instance create --name test1.example.com --size g1.small --region svg1 --ssh-key default",
	Long:    `Create a new instance with the described specification under your current account`,
	Run: func(cmd *cobra.Command, args []string) {
		if instanceCreateName == "" {
			instanceCreateName = api.InstanceSuggestHostname()
		}

		params := api.InstanceParams{
			Name:        instanceCreateName,
			Size:        instanceCreateSize,
			Region:      instanceCreateRegion,
			SSHKey:      instanceCreateSSHKey,
			Template:    instanceCreateTemplate,
			InitialUser: instanceCreateInitialUser,
			PublicIP:    instanceCreatePublicIP,
		}
		_, err := api.InstanceCreate(params)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
	},
}

func init() {
	instanceCmd.AddCommand(instanceCreateCmd)
	instanceCreateCmd.Flags().StringVarP(&instanceCreateName, "name", "n", "", "Name of the account; lowercase, hyphen separated. If you don't specify one, a random one will be used.")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateSize, "size", "s", "g1.small", "The size from 'civo sizes'")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateRegion, "region", "r", "svg1", "The region from 'civo regions'")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateSSHKey, "ssh-key", "k", "default", "The SSH key name from 'civo sshkeys'")
	instanceCreateCmd.Flags().BoolVarP(&instanceCreatePublicIP, "public-ip", "p", true, "Should a public IP address be allocated")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateTemplate, "template", "t", "ubuntu-14.04", "The template from 'civo templates'")
	instanceCreateCmd.Flags().StringVarP(&instanceCreateInitialUser, "initial-user", "u", "civo", "The default user to create")
}
