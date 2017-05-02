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

	"github.com/spf13/cobra"
	"io/ioutil"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"reflect"
)

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "stack commands",
	Long: `Commands to add or remove a service from your stack`,
}
// stackCmd represents the stack command
var stackAddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// get the current stack file
		stackFile, err := cmd.Flags().GetString("stack-file")
		if (stackFile == "") {
			stackFile = viper.GetString("data_dir") + "/dm.yml"
		}
		dat, err := ioutil.ReadFile(stackFile)
		if (err != nil) {
			fmt.Println("Could not read stackfile: " + stackFile)
			return
		}

		mp := make(map[interface{}]interface{})

		err = yaml.Unmarshal([]byte(dat), &mp)

		//fmt.Print(string(dat))
		//err = yaml.Unmarshal([]byte(dat), &f)
		//if (err != nil) {
		//	fmt.Println("Could not unmarshall stackfile data")
		//	return
		//}

		//find the services key, and see if a service with this name exists
		printMap(mp, 0)

		if str, ok := mp["services"].(map[string]interface{}); ok {
			fmt.Print(str)
		} else {
			fmt.Println("not string interface map")
		}

		//if yes then exit

		//if no then add the service to the parsed array

		//save the file
	},
}

func printMap(m map[interface{}]interface{}, indent int) {
	for k, v := range m {
		ind := generateIndent(indent)
		switch vv := v.(type) {
		case string:
			fmt.Println(ind, k, "is string: ", vv)
		case int:
			fmt.Println(ind, k, "is int: ", vv)
		case []interface{}:
			fmt.Println(ind, k, "is a yaml list: ")
			for i, u := range vv {
				fmt.Println(ind + "  ", i, u)
			}
		case map[string]interface{}:
			fmt.Println(ind, k, "is a yaml map")
			printMap(map[interface{}]interface{}(vv), indent + 1)
		case map[interface{}]interface{}:
			printMap(vv, indent + 1)
		default:
			t := reflect.TypeOf(vv)
			fmt.Println(ind, k, "is of a type I don't know how to handle: ", t)
		}
	}
}

func generateIndent(indent int) string {
	var output string
	for i := 0; i < indent; i++ {
		output = output + " "
	}
	return output

}

var stackRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("stack called")
	},
}

func init() {
	RootCmd.AddCommand(stackCmd)
	stackCmd.AddCommand(stackAddCmd)
	stackCmd.AddCommand(stackRmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	stackCmd.PersistentFlags().String("stack-file", "", "The path and name of the stack file to edit")
	stackCmd.PersistentFlags().String("yaml", "", "single line of yaml to add to your stack")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
