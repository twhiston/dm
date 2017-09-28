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
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the local environment",
	Long:  `Stop the local environment`,
	Run: func(cmd *cobra.Command, args []string) {

		deleteLockFile()
		fmt.Println("---> Stopping dm containers")
		RunScript("/bin/sh", "-c", "docker-compose -f "+viper.GetString("data_dir")+"/dm.yml stop")
		listContainers()
	},
}

func init() {
	RootCmd.AddCommand(stopCmd)

}
