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

	"errors"
	"github.com/matishsiao/goInfo"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// requirementsCmd represents the requirements command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Test for all requirements",
	Long: `Running this without a subcommand will execute all the requirements checks`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, element := range cmd.Commands() {
			// element is the element from someSlice for where we are
			err := element.RunE(cmd, args)
			if err != nil {
				fmt.Println("Requirement failed to be met")
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
		fmt.Println("all requirements met")
	},
}

func init() {
	checkCmd.AddCommand(osxReqCmd)
	checkCmd.AddCommand(dockerReqCmd)
	checkCmd.AddCommand(socatReqCmd)
	checkCmd.AddCommand(apacheReqCmd)

	RootCmd.AddCommand(checkCmd)

}

var osxReqCmd = &cobra.Command{
	Use:   "osx",
	Short: "Test for osx requirement",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		gi := goInfo.GetInfo()
		if gi.GoOS != "darwin" {
			return errors.New("This Installer cannot be run without OSX")
		}
		return nil
	},
}

var dockerReqCmd = &cobra.Command{
	Use:   "docker",
	Short: "Test for docker requirement",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := exec.Command("docker").Run(); err != nil {
			return errors.New("	---> Could not find docker on your system\n	     Please install Docker for Mac before running this program\n 	     https://docs.docker.com/docker-for-mac")
		}
		_, exists := os.LookupEnv("DOCKER_HOST")
		if exists {
			return errors.New("	Found something checking for docker envs.\n	This suggests you have the old docker toolbox, please install docker for mac and unset docker vars\n 	See: https://docs.docker.com/docker-for-mac/docker-toolbox/#/setting-up-to-run-docker-for-mac")
		}
		return nil
	},
}

var socatReqCmd = &cobra.Command{
	Use:   "socat",
	Short: "Test for socat requirement",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := exec.Command("socat", "-V").Run(); err != nil {
			return errors.New("	---> Could not find socat on your system\n		Try `brew install socat`")
		}
		return nil
	},
}

var apacheReqCmd = &cobra.Command{
	Use:   "apache",
	Short: "Test apache is off",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := exec.Command("apachectl", "stop").Run(); err != nil {
			return errors.New("	---> Could not stop apache, try again with\n sudo dm check")
		}
		return nil
	},
}
