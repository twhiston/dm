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
	"os"
	"path"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Helper commands,",
	Long:  `For help using the command itself use the --help or -h flags`,
}

var helpVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "dm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("version"))
	},
}

// pwdCmd represents the pwd command
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Find where the dm executable is located",
	Run: func(cmd *cobra.Command, args []string) {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := path.Dir(ex)
		fmt.Println(exPath)
	},
}

func init() {
	RootCmd.AddCommand(helpCmd)

	helpCmd.AddCommand(helpVersionCmd)

	helpCmd.AddCommand(pwdCmd)

}
