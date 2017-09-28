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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os/exec"
)

// startXdebugCmd represents the startXdebug command
var startXdebugCmd = &cobra.Command{
	Use:   "xdebug",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---> Setup Xdebug Loopbacks")
		fmt.Println("---> Ensuring loopback to 10.254.254.254 exists for xdebug listeners")
		RunScript("/bin/sh", "-c", "sudo ifconfig lo0 alias 10.254.254.254")
		fmt.Println("---> Running socat for phpstorm docker integration on http://127.0.0.1:2376")
		command := exec.Command("/bin/sh", "-c", "socat TCP-LISTEN:2376,reuseaddr,fork,bind=127.0.0.1 UNIX-CLIENT:/var/run/docker.sock")
		command.Start()
		viper.Set("init.xdebug", true)
	},
}

func init() {
	startCmd.AddCommand(startXdebugCmd)
}
