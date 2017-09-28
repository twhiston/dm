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
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "add or remove env variables",
	Long:  `Easily add or remove environment variables`,
}

func init() {
	RootCmd.AddCommand(envCmd)

	envCmd.PersistentFlags().String("env-var", "", "The ENV variable name to add")
	envCmd.PersistentFlags().String("file", "", "Name of the user profile file")
	envCmd.PersistentFlags().String("value", "", "The ENV variable value to add")
}

func saveUseProfileFile(envFile string) {
	viper.Set("envfile", envFile)
	saveConfig(viper.AllSettings(), getConfigPath()+"/"+getConfigFileName())
	fmt.Println('\n', "---> User profile file update")
	fmt.Println("---> Your preferred user profile file has been updated to `" + envFile + "`")
	fmt.Println("---> Your choice has been saved for later usage", '\n')
}

func getenvFile() (string, string) {
	envFile, _ := envCmd.PersistentFlags().GetString("file")
	if envFile == "" {
		envFile = viper.GetString("envfile")
		if envFile == "" {
			envFile = userHomeDir() + "/.bash_profile"
			saveUseProfileFile(envFile)
		}
	} else if envFile != viper.GetString("envfile") {
		saveUseProfileFile(envFile)
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
	return strings.Contains(envFileString, "export "+envVarName+"=")
}

func getenvstring(envVarName string, envVarValue string) string {
	return "export " + envVarName + "=" + envVarValue
}
