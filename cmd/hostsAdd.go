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
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a hosts file entry",
	Long:  `Add a site to the local hosts file entry`,
	Run: func(cmd *cobra.Command, args []string) {
		// read the whole file at once
		hostFile, hostFileString := getHostsFile()

		if len(args) <= 0 {
			fmt.Println("Must have host name as argument")
			return
		}
		hostName := args[0]

		if hostExists(hostName, hostFileString) {
			fmt.Println("host already exists in hosts file")
			return
		}

		//build full string
		hostString := getHostString(hostName)

		//write string to file
		err := appendStringToFile(hostFile, "\n"+hostString)
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
