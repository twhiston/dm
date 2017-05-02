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
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleanup your environment",
	Long: `Other commands may optionally implement a clean subcommand, which will do something to clean up your environment for you.
	Running the base clean command will execute all of the clean steps, similarly to start, run with --help or -h to see subcommands`,
	Run: func(cmd *cobra.Command, args []string) {

		for _, element := range cmd.Commands() {
			// element is the element from someSlice for where we are
			element.Run(cmd, args)
		}
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}
