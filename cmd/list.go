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

	"bytes"
	"github.com/spf13/cobra"
	"os/exec"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List running docker containers",
	Long: `Exactly analogous to docker ps,
normally just called by other commands to show containers during operations`,
	Run: func(cmd *cobra.Command, args []string) {
		listContainers()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func listContainers() {
	fmt.Println("---> Listing Running Docker Containers")
	command := exec.Command("docker", "ps")
	var out bytes.Buffer
	command.Stdout = &out
	command.Run()
	fmt.Print(out.String())
}
