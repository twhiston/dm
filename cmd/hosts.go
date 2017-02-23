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
	"os"
	"io/ioutil"
	"strings"
)

// hostsCmd represents the hosts command
var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "add or remove hosts file entries",
	Long: `Easily add or remove dev site aliases from your local hosts file here

	This allows you easily manage redirects via the nginx proxy container`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use `dm hosts --help` for commands")
	},
}

func init() {
	RootCmd.AddCommand(hostsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostsCmd.PersistentFlags().String("foo", "", "A help for foo")
	hostsCmd.PersistentFlags().String("host", "", "The hostname to add")
	hostsCmd.PersistentFlags().String("file", "/private/etc/hosts", "Full path to hostsfile")
	hostsCmd.PersistentFlags().String("dhost", "0.0.0.0", "IP of the docker host")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

/**
 * Append string to the end of file
 *
 * path: the path of the file
 * text: the string to be appended. If you want to append text line to file,
 *       put a newline '\n' at the of the string.
 */
func appendStringToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func getHostsFile() (string, string){
	hostFile, err := hostsCmd.PersistentFlags().GetString("file")
	hostFileData, err := ioutil.ReadFile(hostFile)
	if err != nil {
		fmt.Println("Failed to open hosts file. You may need to run with sudo")
		os.Exit(1)
	}
	return hostFile, string(hostFileData)
}

func hostExists(hostName string, hostFileString string) bool{
	//check if host exists
	return strings.Contains(hostFileString, hostName)
}

func getHostString(hostName string) string {
	dHost, _ := hostsCmd.PersistentFlags().GetString("dhost")
	return dHost + " " + hostName
}
