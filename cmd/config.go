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
	"os"
	"runtime"
	"strings"
)

var configCreate bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage dm local config file",
	Long:  ``,
}

func init() {
	RootCmd.AddCommand(configCmd)
}

func getConfigPath() string {
	configPath := userHomeDir() + "/.dm"
	if strings.HasPrefix(configPath, "$HOME") {
		configPath = userHomeDir() + configPath[5:]
	}
	return configPath
}

func getConfigFileName() string {
	return "config.yml"
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func createConfigDir(configPath string) string {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Creating data dir: " + configPath)
		err = os.Mkdir(configPath, 0755)
		if err != nil {
			fmt.Println("Could not create config directory:" + configPath + " try running with sudo ")
		}
	}

	return configPath
}
