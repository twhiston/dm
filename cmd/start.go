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
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Docker 4 Mac local Environment",
	Long:  `Start the local environment by running all of the Extensions`,
	Run: func(cmd *cobra.Command, args []string) {
		forceFlag, _ := cmd.PersistentFlags().GetBool("force")
		createLockFile(forceFlag)

		for _, element := range cmd.Commands() {
			// element is the element from someSlice for where we are
			element.Run(cmd, args)
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().BoolP("force", "f", false, "Force the start command to run even if a lock file exists")
	//startCmd.PersistentFlags().BoolP("strict", "s", true, "Stop running if requirements are not met")

}
