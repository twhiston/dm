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
	if //noinspection ALL
	runtime.GOOS == "windows" {
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
