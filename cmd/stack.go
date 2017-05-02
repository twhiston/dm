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
)

// containersCmd represents the containers command
var stackCmd = &cobra.Command{
	Use:   "z_stack",
	Short: "start the containers only",
	Long:  `The weird command name ensure that this gets sorted last in the child commands :(`,
	Run: func(cmd *cobra.Command, args []string) {
		//docker compose up
		fmt.Println("---> Starting dm containers")
		RunScript("/bin/sh", "-c", "docker-compose -f "+viper.GetString("data_dir")+"/dm.yml up -d")
		listContainers()
	},
}

func init() {
	startCmd.AddCommand(stackCmd)
}
