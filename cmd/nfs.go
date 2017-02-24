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

	"github.com/libgit2/git2go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twhiston/dm/process"
	"os"
)

// startNfsCmd represents the startNfs command
var startNfsCmd = &cobra.Command{
	Use:   "startNfs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("	---> Start NFS shares")

		if !viper.GetBool("nfsInit") {
			fmt.Println("Nfs not initialized, run init")
			os.Exit(1)
		}
		//Run the command
		nfsDir := viper.GetString("data_dir") + "/nfs"
		process.RunScript(nfsDir + "/d4m-nfs.sh")
	},
}

var initNfsCmd = &cobra.Command{
	Use:   "startNfs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("	---> Setup NFS shares")
		nfsDir := viper.GetString("data_dir") + "/nfs"
		fmt.Println(nfsDir)
		if _, err := os.Stat(nfsDir); os.IsNotExist(err) {
			//If the directory doesn't exist then make it and clone the helper repo we are using
			fmt.Print("Creating nfs mount script dir: ")
			fmt.Println(nfsDir)
			process.HandleError(os.Mkdir(nfsDir, 0755))
			_, err = git.Clone("https://github.com/IFSight/d4m-nfs", nfsDir, &git.CloneOptions{})
			process.HandleError(err)
			//Now the repo is cloned copy in our unique assets
			//Get the data from the config file
			//Turn it into a text file and write it to the /etc/folder
			//data := getAsset("nfs/d4m-nfs-mounts.txt")
			//writeAsset(nfsDir+"/etc/d4m-nfs-mounts.txt", data)
		}
		viper.Set("nfsInit", true)
	},
}

func init() {
	startCmd.AddCommand(startNfsCmd)
	initCmd.AddCommand(initNfsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startNfsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startNfsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
