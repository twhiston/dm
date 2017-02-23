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
	"github.com/spf13/viper"
	"os"
	"strings"
	"gopkg.in/yaml.v2"
)


var configRst bool
var configCreate bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("config called")
		if configRst == true {
			fmt.Println("Reset config file to defaults")
			os.Remove(getConfigPath()+"/dm.yml")
			configCreate = true
		}
		if configCreate == true {
			createConfig()
		}
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")
	configCmd.Flags().BoolVarP(&configRst,"reset", "r", false, "Reset config file")
	configCmd.Flags().BoolVarP(&configCreate,"create", "c", false, "Create config file")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func getConfigPath() string {
	configPath := "/Users/Shared/.dm"
	if strings.HasPrefix(configPath, "$HOME") {
		configPath = userHomeDir() + configPath[5:]
	}
	return configPath
}

func createConfig() error {


	cfgpath := getConfigPath() + "/config.yml"
	b, err := yaml.Marshal(viper.AllSettings())
	if err != nil {
		return err
	}

	f, err := os.Create(cfgpath)
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(string(b))

	return nil
}