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
	"os"
	"strings"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "add or remove env variables",
	Long: `Easily add or remove environement variables

	This allows you to easily configure Blackfire environment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use `dm env --help` for commands")
	},
}

func init() {
	RootCmd.AddCommand(envCmd)

	envCmd.PersistentFlags().String("variable", "", "The ENV variable name to add")
	envCmd.PersistentFlags().String("file", ".bash_profile", "Name of the user profile file")
	envCmd.PersistentFlags().String("value", "", "The ENV variable value to add")

}

func getenvFile() (string, string) {
	envFile, err := envCmd.PersistentFlags().GetString("file")
	if envFile == ".bash_profile" {
		envFile = userHomeDir() + "/" + envFile
	}
	envFileData, err := ioutil.ReadFile(envFile)
	if err != nil {
		fmt.Println("Failed to open user profile file. You may need to run with sudo")
		os.Exit(1)
	}
	return envFile, string(envFileData)
}

func envExists(envVarName string, envFileString string) bool {
	//check if env variable exists
	return strings.Contains(envFileString, "export " + envVarName + "=")
}

func getenvstring(envVarName string, envVarValue string) string {
	return "export " + envVarName + "=" + envVarValue
}
