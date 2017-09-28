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
	"gopkg.in/yaml.v2"
)

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "do things to the stack",
	Long:  ``,
}

// containersCmd represents the containers command
var stackStartCmd = &cobra.Command{
	Use:   "z_stack",
	Short: "start the containers only",
	Long:  `The weird command name ensure that this gets sorted last in the child commands :(`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---> Starting dm containers")
		RunScript("/bin/sh", "-c", "docker-compose -f "+viper.GetString("data_dir")+"/dm.yml up -d")
		listContainers()
	},
}

var cleanStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "clean stack",
	Long:  `delete all containers in the stack file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---> Deleting dm stack containers")
		RunScript("/bin/sh", "-c", "docker-compose -f "+viper.GetString("data_dir")+"/dm.yml rm -f")

	},
}

// stackStartCmd represents the stack command
var stackAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add elements to the stack",
	Long: `It is possible to add multiple items to the stack at once and they will be processed in the order

	snippet
	file


SNIPPETS

Strings passed to the --snippet flag should contain only the definition of the service they define, at the top level.
For Example:

dm stack add redis --snippet='image: redis
container_name: redis
ports:
- 6397:6379
networks:
- bridge
environment:
- VIRTUAL_HOST=redis.dev'

as the service name will be derived from the argument

FILE

Strings passed to the --file command should be the full path to a docker-compose file to merge with the current stack.
The file should be in the docker-compose 2 format, and services are expected to be found beneath the top level services key.
If the current stack has a service with the same key as a service in the merge file it will be REPLACED`,
	Run: func(cmd *cobra.Command, args []string) {

		stackFile, err := getActiveStackPath(cmd)
		HandleError(err, false)
		stack, filepath, err := getStack(stackFile)
		HandleError(err, false)
		services := getServices(stack)

		if len(args) > 0 {
			sName := args[0]

			if hasService(sName, services) {
				fmt.Println("Service already exists, remove it first with \"dm stack rm", sName+"\"")
				return
			}

			snippet, err := cmd.Flags().GetString("snippet")
			HandleError(err, false)
			if snippet != "" {

				//parse it as yml
				var parsed interface{}
				yaml.Unmarshal([]byte(snippet), &parsed)

				services[sName] = parsed
			}
		}

		merge, err := cmd.Flags().GetString("merge")
		if merge != "" {
			mergeFile, _, err := getStack(merge)
			HandleError(err, false)
			if mergeFile.GetString("version") != "2" {
				fmt.Println("File merges only support version 2 docker files")
			}

			softMerge, err := cmd.Flags().GetBool("merge-soft")
			HandleError(err, true)
			mergeServices := getServices(mergeFile)
			for k, v := range mergeServices {
				if !softMerge {
					services[k] = v
					continue
				}
				if _, ok := services[k]; !ok {
					services[k] = v
				}

			}

			if len(mergeServices) > 0 {
				fmt.Println("merged services")
			}
		}

		saveConfig(stack.AllSettings(), filepath)
		fmt.Println("Stack file saved")

	},
}

var stackRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a stack item",
	Long:  `Deletes an item from the stack file`,
	Run: func(cmd *cobra.Command, args []string) {
		stackPath, err := getActiveStackPath(cmd)
		HandleError(err, false)
		stack, filepath, err := getStack(stackPath)
		HandleError(err, false)
		services := getServices(stack)
		if !hasService(args[0], services) {
			fmt.Println("service does not exist")
			return
		}

		delete(services, args[0])
		stack.Set("services", services)
		saveConfig(stack.AllSettings(), filepath)
		fmt.Println(args[0], "removed from stack. Restart dm if running")

	},
}

func getActiveStackPath(cmd *cobra.Command) (string, error) {
	stackFile, err := cmd.Flags().GetString("stack")
	if err != nil {
		return "", err
	}
	// get the current stack file
	if stackFile == "" {
		fmt.Println("Getting stack file from config")
		stackFile = viper.GetString("data_dir") + "/dm.yml"
	}
	return stackFile, nil
}

func getStack(stackFile string) (*viper.Viper, string, error) {

	sv := viper.New()
	sv.SetConfigFile(stackFile)
	if err := sv.ReadInConfig(); err == nil {
		fmt.Println("Loading Stack:", sv.ConfigFileUsed())
	} else {
		return nil, stackFile, err
	}

	return sv, stackFile, nil

}

func getServices(v *viper.Viper) map[string]interface{} {
	return v.GetStringMap("services")
}

func hasService(key string, services map[string]interface{}) bool {
	_, ok := services[key]
	return ok
}

func init() {
	RootCmd.AddCommand(stackCmd)
	stackCmd.AddCommand(stackAddCmd)
	stackCmd.AddCommand(stackRmCmd)

	startCmd.AddCommand(stackStartCmd)
	cleanCmd.AddCommand(cleanStackCmd)

	stackCmd.PersistentFlags().String("stack", "", "The path and name of the stack file to edit")
	stackAddCmd.Flags().String("snippet", "", "A snippet that contains the definition of the stack service that you want to have. NOT including the top level services, or service name keys")
	stackAddCmd.Flags().String("merge", "", "A path to a file of services to be merged with the existing stack file.")
	stackAddCmd.Flags().Bool("merge-soft", false, "If true definitions in current stack file will be preferred to merge file")

}
