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
	"os"
	"errors"
	//"bufio"

	"github.com/spf13/cobra"
)

// startBlackfireCmd represents the startXdebug command
var startBlackfireCmd = &cobra.Command{
	Use:   "blackfire",
	Short: "Checks if environment is correctly set for Blackfire",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("	---> Check Blackfire requirement")
		fmt.Println("		---> Ensuring Blackfire environment variables are defined")

		err := checkBlackfireEnv(cmd, args)
		if err != nil {
			fmt.Println(err.Error())
			err := setBlackfireEnv(cmd, args)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	},
}

func checkBlackfireEnv(cmd *cobra.Command, args []string) error {
	if os.Getenv("BLACKFIRE_SERVER_ID") == "" {
		return errors.New("\n		/!\\ WARNING /!\\\n		It seems that your environment is not set properly\n		To permanently fix this issue, add the following block in your ~/.bashrc or ~/.bash_profile :\n		\n		### Start Blackfire configuration\n		export BLACKFIRE_SERVER_ID=<YOUR_SERVER_ID>\n		export BLACKFIRE_SERVER_TOKEN=<YOUR_SERVER_TOKEN>\n		export BLACKFIRE_CLIENT_ID=<YOUR_CLIENT_ID>\n		export BLACKFIRE_CLIENT_TOKEN=<YOUR_CLIENT_TOKEN>\n		### End Blackfire configuration\n		\n		and then log back in\n		\n		You can temporarily fix it now:\n		")
	}
	return nil
}

func setBlackfireEnv(cmd *cobra.Command, args []string) error {
	fmt.Print("Enter your SERVER_ID: ")
    var server_id string
    fmt.Scanln(&server_id)

    fmt.Print("Enter your SERVER_TOKEN: ")
    var server_token string
    fmt.Scanln(&server_token)

    fmt.Print("Enter your CLIENT_ID: ")
    var client_id string
    fmt.Scanln(&client_id)

    fmt.Print("Enter your CLIENT_TOKEN: ")
    var client_token string
    fmt.Scanln(&client_token)

	if server_id == "" || server_token == "" || client_id == "" || client_token == "" {
		return errors.New("\n		/!\\ ERROR /!\\\n		You did not enter a valid value\n")
	}

	os.Setenv("BLACKFIRE_SERVER_ID", server_id)
	os.Setenv("BLACKFIRE_SERVER_TOKEN", server_token)
	os.Setenv("BLACKFIRE_CLIENT_ID", client_id)
	os.Setenv("BLACKFIRE_CLIENT_TOKEN", client_token)

	return nil
}

func init() {
	startCmd.AddCommand(startBlackfireCmd)
}
