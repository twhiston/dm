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
	"io/ioutil"
	"regexp"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a hosts file entry",
	Long: `Remove a site from the local hosts file entry`,
	Run: func(cmd *cobra.Command, args []string) {

		hostFile, hostFileString := getHostsFile()
		hostName, _ := hostsCmd.PersistentFlags().GetString("host")
		if hostName == "" {
			fmt.Println("Must have --host parameter")
			return
		}

		if(!hostExists(hostName, hostFileString)){
			fmt.Println("host does not exist in hosts file")
			return
		}

		hostString := getHostString(hostName)

		re := regexp.MustCompile(hostString)
		res := re.ReplaceAllString(hostFileString, "")

		err := ioutil.WriteFile(hostFile, []byte(res), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("Removed " + hostString + " from " + hostFile)
	},
}

func init() {
	hostsCmd.AddCommand(rmCmd)
}
