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
