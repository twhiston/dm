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
	"path"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Helper commands,",
	Long:  `For help usng the command itself use the --help or -h flags`,
}

var helpVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "dm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("version"))
	},
}

// pwdCmd represents the pwd command
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Find where the dm executable is located",
	Run: func(cmd *cobra.Command, args []string) {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := path.Dir(ex)
		fmt.Println(exPath)
	},
}

func init() {
	RootCmd.AddCommand(helpCmd)

	helpCmd.AddCommand(helpVersionCmd)

	helpCmd.AddCommand(pwdCmd)

}
