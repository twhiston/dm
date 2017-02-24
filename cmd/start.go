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
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Docker 4 Mac local Environment",
	Long: `Start the local environment by running all of the Extensions`,
	Run: func(cmd *cobra.Command, args []string) {
		for _,element := range cmd.Commands() {
			// element is the element from someSlice for where we are
			element.Run(cmd, args);
		}
		//cmd.Commands()
		//o := SetUpListeners()
		//strict, _ := cmd.PersistentFlags().GetBool("strict")
		//o.Trigger("check-requirements", strict)
		//forceFlag, _ := cmd.PersistentFlags().GetBool("force")
		//createLockFile(cfgFilePath, forceFlag)
		//o.Trigger("start", cfgFilePath)
	},
}

func init() {
	RootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().BoolP("force", "f", false, "Force the start command to run even if a lock file exists")
	startCmd.PersistentFlags().BoolP("strict", "s", true, "Stop running if requirements are not met")

}


