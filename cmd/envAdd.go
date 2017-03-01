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
