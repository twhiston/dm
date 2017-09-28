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

// mariadbCmd represents the mariadb command
var mariadbInitCmd = &cobra.Command{
	Use:   "mariadb",
	Short: "initialize mariadb",
	Long: `This command will create all the necessary data directories for mariadb.
	It will only do this if the directories do not exist already, so is safe to execute as part of the base init command
	once data has been stored`,
	Run: func(cmd *cobra.Command, args []string) {
		var mariaDir = getMariaDir()
		if _, err := os.Stat(mariaDir); os.IsNotExist(err) {
			//Install stuff if it doesnt exist
			fmt.Println("---> Setup Mariadb")
			fmt.Print("---> Creating mariadb dir: ")
			fmt.Println(mariaDir)
			os.Mkdir(mariaDir, 0777)

			fmt.Print("---> Creating mariadb data dir: ")
			var mariaDataDir = mariaDir + "/data"
			fmt.Println(mariaDataDir)
			os.Mkdir(mariaDataDir, 0777)

			fmt.Println("---> Copying environment assets")
			data := GetAsset("database/.env")
			WriteAsset(mariaDir+"/.env", data)
		}
		viper.Set("init.mariadb", true)
		fmt.Println("---> MariaDb Initialized <--- ")
	},
}

func getMariaDir() string {
	return viper.GetString("data_dir") + "/maria"
}

func init() {
	initCmd.AddCommand(mariadbInitCmd)
}
