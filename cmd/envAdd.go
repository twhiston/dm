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

// envAddCmd represents the add command
var setenvCmd = &cobra.Command{
	Use:   "add",
	Short: "add a env file entry",
	Long:  `Add a site to the local env file entry`,
	Run: func(cmd *cobra.Command, args []string) {
		// read the whole file at once
		envFile, envFileString := getenvFile()
		envVarName, _ := envCmd.PersistentFlags().GetString("env-var")
		if envVarName == "" {
			fmt.Println("--env-var cannot be blank")
			return
		}

		envVarValue, _ := envCmd.PersistentFlags().GetString("value")
		if envVarValue == "" {
			fmt.Println("--value cannot be blank")
			return
		}

		if envExists(envVarName, envFileString) {
			fmt.Println("ENV variable already exists in user profile file")
			return
		}

		//build full string
		envstring := getenvstring(envVarName, envVarValue)

		//write string to file
		err := appendStringToFile(envFile, "\n"+envstring)
		if err != nil {
			fmt.Println("Failed to add env to env file. You may need to run with sudo")
			return
		}

		fmt.Println("Added " + envstring + " to " + envFile)
	},
}

func init() {
	envCmd.AddCommand(setenvCmd)
}
