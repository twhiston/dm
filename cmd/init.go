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
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your local environment for dm",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		dir := viper.GetString("data_dir")

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			//Install stuff if it doesnt exist
			fmt.Print("---> Creating data dir: ")
			fmt.Println(dir)
			os.Mkdir(dir, 0777)
		}

		data := GetAsset("dm.yml")
		WriteAsset(dir+"/dm.yml", data)

		for _, element := range cmd.Commands() {
			// element is the element from someSlice for where we are
			element.Run(cmd, args)
		}
		saveConfig(viper.AllSettings(), getConfigPath()+"/"+getConfigFileName())
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().String("sharepath", "/Users/Shared/.dm", "Share path for sites")
}
