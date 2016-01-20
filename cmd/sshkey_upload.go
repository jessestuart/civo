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
	"io/ioutil"
	"os"
	"strings"

	"github.com/absolutedevops/civo/api"
	"github.com/spf13/cobra"
)

var sshKeyUploadName string
var sshKeyUploadPublicKeyFile string

var sshKeyUploadCmd = &cobra.Command{
	Use:     "upload",
	Aliases: []string{"new", "store"},
	Short:   "Upload an SSH key",
	Example: "civo sshkey upload --name default --public-key ~/.ssh/id_rsa.pub",
	Long:    `Upload an SSH public key, then this can be chosen when creating an instance`,
	Run: func(cmd *cobra.Command, args []string) {
		if sshKeyUploadName == "" {
			fmt.Println("You need to specify a name with --name in order to upload a key")
			os.Exit(-3)
		}
		if sshKeyUploadPublicKeyFile == "" {
			fmt.Println("You need to specify the filename of a valid SSH Public Key with --public-key in order to upload a key")
			os.Exit(-3)
		}

		params := api.SshKeyParams{
			Name: sshKeyUploadName,
		}
		filename := strings.Replace(sshKeyUploadPublicKeyFile, "~", "$HOME", -1)
		filename = os.ExpandEnv(filename)
		buf, err := ioutil.ReadFile(filename)
		if err == nil {
			params.PublicKey = string(buf)
		}

		_, err = api.SshKeyCreate(params)
		if err != nil {
			fmt.Printf("An error occured: ", err)
			return
		}
		fmt.Printf("Uploaded SSH key called `%s`\n", sshKeyUploadName)
	},
}

func init() {
	sshKeyCommand.AddCommand(sshKeyUploadCmd)
	sshKeyUploadCmd.Flags().StringVarP(&sshKeyUploadName, "name", "n", "", "A nice name to be used for the key; lowercase and hyphen separated")
	sshKeyUploadCmd.Flags().StringVarP(&sshKeyUploadPublicKeyFile, "public-key", "k", "", "The filename of the public key file, e.g. ~/.ssh/id_rsa.pub")
}
