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
)

// startBlackfireCmd represents the startXdebug command
var startBlackfireCmd = &cobra.Command{
	Use:   "blackfire",
	Short: "Checks if environment is correctly set for Blackfire",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("	---> Blackfire recommendations")
		fmt.Println("	---> If you plan to use Blackfire, make sure that:")
		fmt.Println("		- Xdebug is not activated in your application container")
		fmt.Println("		- Blackfire Probe is installed in your application container")
		fmt.Println("		- The following command passes successfully:")
		fmt.Println("			 dm check blackfire")
	},
}

func init() {
	startCmd.AddCommand(startBlackfireCmd)
}
