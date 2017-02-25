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
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your local environment for dm",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		dir := viper.GetString("data_dir")

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			//Install stuff if it doesnt exist
			fmt.Print("		---> Creating data dir: ")
			fmt.Println(dir)
			os.Mkdir(dir, 0777)
		}

		data := GetAsset("dm.yml")
		WriteAsset(dir+"/dm.yml", data)

		for _, element := range cmd.Commands() {
			// element is the element from someSlice for where we are
			element.Run(cmd, args)
		}
		saveConfig()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().String("sharepath", "/Users/Shared/.dm", "Share path for sites")
}
