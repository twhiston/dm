// Copyleft © 2016 Tom Whiston <tom.whiston@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

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
	Long:  `Running this without a subcommand will execute all the requirements checks`,
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
	checkCmd.AddCommand(blackfireReqCmd)

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
			return errors.New("---> Could not find docker on your system\nPlease install Docker for Mac before running this program\n 	     https://docs.docker.com/docker-for-mac")
		}
		_, exists := os.LookupEnv("DOCKER_HOST")
		if exists {
			return errors.New("Found something checking for docker envs.\nThis suggests you have the old docker toolbox, please install docker for mac and unset docker vars\n 	See: https://docs.docker.com/docker-for-mac/docker-toolbox/#/setting-up-to-run-docker-for-mac")
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
			return errors.New("---> Could not find socat on your system\nTry `brew install socat`")
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
			return errors.New("---> Could not stop apache, try again with\n sudo dm check")
		}
		return nil
	},
}

var blackfireReqCmd = &cobra.Command{
	Use:   "blackfire",
	Short: "Test if environment is correctly set for Blackfire",
	Long:  `Will check if your environment variables contains BLACKFIRE_SERVER_ID`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("---> Blackfire recommendations")
		fmt.Println("---> If you plan to use Blackfire, make sure that:")
		fmt.Println("	- Xdebug is not activated in your application container")
		fmt.Println("	- Blackfire Probe is installed in your application container")
		if os.Getenv("BLACKFIRE_SERVER_ID") == "" || os.Getenv("BLACKFIRE_SERVER_TOKEN") == "" {
			return errors.New(`		/!\\ ERROR /!\\
		It seems that your environment is not set properly
		First, make sure you are not using this command with 'sudo'.
		To fix this issue, run the following commands:

		dm env add --variable=BLACKFIRE_SERVER_ID --value=<YOUR_SERVER_ID>
		dm env add --variable=BLACKFIRE_SERVER_TOKEN --value=<YOUR_SERVER_TOKEN>
		dm env add --variable=BLACKFIRE_CLIENT_ID --value=<YOUR_CLIENT_ID>
		dm env add --variable=BLACKFIRE_CLIENT_TOKEN --value=<YOUR_CLIENT_TOKEN>

		And then log back in

		Note:
			if you want to store your environment variables in a different file than '.bash_profile'
			you can add the following flag to the commands above:
			--file=<FULL_PATH_OF_YOUR_PREFERRED_FILE>`)
		}
		return nil
	},
}
