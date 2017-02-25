// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

// mariadbCmd represents the mariadb command
var mariadbInitCmd = &cobra.Command{
	Use:   "mariadb",
	Short: "initialize mariadb",
	Long: `This command will create all the necessary data directories for mariadb.
	It will only do this if the directories do not exist already, so is safe to execute as part of the base init command
	once data has been stored`,
	Run: func(cmd *cobra.Command, args []string) {
		var mariaDir = getMariaDir()
		if _, err := os.Stat(mariaDir); os.IsNotExist(err) {
			//Install stuff if it doesnt exist
			fmt.Println("	---> Setup Mariadb")
			fmt.Print("		---> Creating mariadb dir: ")
			fmt.Println(mariaDir)
			os.Mkdir(mariaDir, 0777)

			fmt.Print("		---> Creating mariadb data dir: ")
			var mariaDataDir = mariaDir + "/data"
			fmt.Println(mariaDataDir)
			os.Mkdir(mariaDataDir, 0777)

			fmt.Println("		---> Copying environment assets")
			data := GetAsset("database/.env")
			WriteAsset(mariaDir+"/.env", data)
		}
		viper.Set("init.mariadb", true)
		fmt.Println("	---> MariaDb Initialized <--- ")
	},
}

func getMariaDir() string {
	return viper.GetString("data_dir") + "/maria"
}

func init() {
	initCmd.AddCommand(mariadbInitCmd)
}
