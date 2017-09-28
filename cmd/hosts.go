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
	"os"
	"strings"
)

// hostsCmd represents the hosts command
var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "add or remove hosts file entries",
	Long: `Easily add or remove dev site aliases from your local hosts file here
This allows you easily manage redirects via the nginx proxy container`,
}

var hostsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list hosts file entries",
	Run: func(cmd *cobra.Command, args []string) {
		_, hostFileString := getHostsFile()
		fmt.Print(hostFileString)
	},
}

func init() {
	RootCmd.AddCommand(hostsCmd)

	hostsCmd.PersistentFlags().String("file", "/private/etc/hosts", "Full path to hostsfile")
	hostsCmd.PersistentFlags().String("dhost", "0.0.0.0", "IP of the docker host")

	hostsCmd.AddCommand(hostsListCmd)

}

/**
 * Append string to the end of file
 *
 * path: the path of the file
 * text: the string to be appended. If you want to append text line to file,
 *       put a newline '\n' at the of the string.
 */
func appendStringToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	return err
}

func getHostsFile() (string, string) {
	hostFile, _ := hostsCmd.PersistentFlags().GetString("file")
	hostFileData, err := ioutil.ReadFile(hostFile)
	if err != nil {
		fmt.Println("Failed to open hosts file. You may need to run with sudo")
		os.Exit(1)
	}
	return hostFile, string(hostFileData)
}

func hostExists(hostName string, hostFileString string) bool {
	//check if host exists
	return strings.Contains(hostFileString, hostName)
}

func getHostString(hostName string) string {
	dHost, _ := hostsCmd.PersistentFlags().GetString("dhost")
	return dHost + " " + hostName
}
