// Copyleft © 2016 Tom Whiston <tom.whiston@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

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
	Long: `Remove a site from the local hosts file entry.
	dHost + " " + hostName forms a regexp which must compile and be found in the hosts file to be removed`,
	Run: func(cmd *cobra.Command, args []string) {

		hostFile, hostFileString := getHostsFile()
		if len(args) <= 0 {
			fmt.Println("Must have host name as argument")
			return
		}
		hostName := args[0]
		if hostName == "" {
			fmt.Println("Must have --host parameter")
			return
		}

		if !hostExists(hostName, hostFileString) {
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
