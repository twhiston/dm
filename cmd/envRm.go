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
	"regexp"
	"io/ioutil"
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
		re := regexp.MustCompile("(?m)[\r\n]+^export "+envVarName+"=.*$")
		res := re.ReplaceAllString(envFileString, "")

		err := ioutil.WriteFile(envFile,[]byte(res), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("Removed " + envVarName + " from " + envFile)
	},
}

func init() {
	envCmd.AddCommand(rmEnvCmd)
}
