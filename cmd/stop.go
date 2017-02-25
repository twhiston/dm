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
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the local environment",
	Long:  `Stop the local environment`,
	Run: func(cmd *cobra.Command, args []string) {

		deleteLockFile()
		fmt.Println("	---> Stopping pxd containers")
		RunScript("/bin/sh", "-c", "docker-compose -f "+viper.GetString("data_dir")+"/dm.yml stop")
		listContainers()
	},
}

func init() {
	RootCmd.AddCommand(stopCmd)

}
