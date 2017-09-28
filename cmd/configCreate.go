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
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
)

// configCreateCmd represents the configCreate command
var configCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a basic config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		saveConfig(viper.AllSettings(), getConfigPath()+"/"+getConfigFileName())
	},
}

func init() {
	configCmd.AddCommand(configCreateCmd)
}

func saveConfig(settings map[string]interface{}, cfgpath string) {

	configPath := getConfigPath()
	createConfigDir(configPath)

	b, err := yaml.Marshal(settings)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(cfgpath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString(string(b))

}
