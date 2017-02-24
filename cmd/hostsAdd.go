// Copyright © 2016 Tom Whiston <tom.whiston@gmail.com>
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

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a hosts file entry",
	Long:  `Add a site to the local hosts file entry`,
	Run: func(cmd *cobra.Command, args []string) {
		// read the whole file at once
		hostFile, hostFileString := getHostsFile()
		hostName, _ := hostsCmd.PersistentFlags().GetString("host")

		if hostExists(hostName, hostFileString) {
			fmt.Println("host already exists in hosts file")
			return
		}

		//build full string
		hostString := getHostString(hostName)

		//write string to file
		err := appendStringToFile(hostFile, "\r\n"+hostString)
		if err != nil {
			fmt.Println("Failed to add host to hosts file. You may need to run with sudo")
			return
		}

		fmt.Println("Added " + hostString + " to " + hostFile)
	},
}

func init() {
	hostsCmd.AddCommand(addCmd)
}
