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

// envAddCmd represents the add command
var rmEnvCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove an env file entry",
	Long:  `Remove a local environment variable`,
	Run: func(cmd *cobra.Command, args []string) {
		// read the whole file at once
		envFile, envFileString := getenvFile()
		envVarName, _ := envCmd.PersistentFlags().GetString("env-var")
		if envVarName == "" {
			fmt.Println("--env-var cannot be blank")
			return
		}

		if !envExists(envVarName, envFileString) {
			fmt.Println("ENV variable does not exist in user profile file")
			return
		}

		//Remove env var
		re := regexp.MustCompile("(?m)[\r\n]+^export " + envVarName + "=.*$")
		res := re.ReplaceAllString(envFileString, "")

		err := ioutil.WriteFile(envFile, []byte(res), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("Removed " + envVarName + " from " + envFile)
	},
}

func init() {
	envCmd.AddCommand(rmEnvCmd)
}
