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
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// startNfsCmd represents the startNfs command
var startNfsCmd = &cobra.Command{
	Use:   "nfs",
	Short: "start nfs sharing in docker machine",
	Long: `This command alters your /etc/exports file if necessary with your nfs sharing configuration
	It has a wait period at the end as it was found that without it sometimes launching containers would fail`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---> Start NFS shares")

		if !viper.GetBool("init.nfs") {
			fmt.Println("---> Nfs not initialized, run init first")
			os.Exit(1)
		}
		//Run the command
		nfsDir := viper.GetString("data_dir") + "/nfs"
		RunScript(nfsDir + "/d4m-nfs.sh")
		//fmt.Print(output)
		fmt.Println("---> Wait for NFS")
		time.Sleep(10000 * time.Millisecond)
		fmt.Println("---> NFS started <---")
	},
}

var initNfsCmd = &cobra.Command{
	Use:   "nfs",
	Short: "Initialize NFS",
	Long: `This command clones the git repo of
		https://github.com/IFSight/d4m-nfs
	and then configures a mounts file based on the current user and the configured directories`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---> Setup NFS shares")
		nfsDir := viper.GetString("data_dir") + "/nfs"
		fmt.Println(nfsDir)
		//if _, err := os.Stat(nfsDir); os.IsNotExist(err) {
		//If the directory doesn't exist then make it and clone the helper repo we are using
		fmt.Println("---> Creating nfs mount script dir ")
		HandleError(os.Mkdir(nfsDir, 0777), true)
		_, err := git.PlainClone(nfsDir, false, &git.CloneOptions{
			URL:      "https://github.com/IFSight/d4m-nfs",
			Progress: os.Stdout,
		})
		HandleError(err, false)
		//Get the data from the config file
		//Turn it into a text file and write it to the /etc/folder
		data := GetAsset("nfs/d4m-nfs-mounts.txt")
		s := string(data[:])
		//add custom shares
		fmt.Println("---> Adding Custom Shares ")
		s += viper.GetString("data_dir") + ":" + viper.GetString("data_dir") + ":0:0 \n"
		s += viper.GetString("share_dir") + ":" + viper.GetString("share_dir") + ":"
		uid := strings.Trim(viper.GetString("uid"), "\r\n")
		s += uid + ":" + viper.GetString("group") + " \n"
		s = strings.TrimSpace(s)
		s += "\n" //Must end with a blank line or the nfs script does not properly iterate the last value
		data = []byte(s)
		WriteAsset(nfsDir+"/etc/d4m-nfs-mounts.txt", data)
		viper.Set("init.nfs", true)
		fmt.Println("---> nfs initialized <---")
	},
}

// startNfsCmd represents the startNfs command
var cleanNfsCmd = &cobra.Command{
	Use:   "nfs",
	Short: "clean nfs sharing in docker machine",
	Long:  `makes your /etc/exports file blank, if this command fails run it with sudo`,
	Run: func(cmd *cobra.Command, args []string) {

		exportFile, _ := cmd.PersistentFlags().GetString("exports-path")
		fmt.Println("---> Make ", exportFile, " file empty")

		res := ""

		err := ioutil.WriteFile(exportFile, []byte(res), 0644)
		if err != nil {
			fmt.Println("Failed to open exports file. You may need to run with sudo")
			os.Exit(1)
		}

	},
}

func init() {
	startCmd.AddCommand(startNfsCmd)
	initCmd.AddCommand(initNfsCmd)
	cleanCmd.AddCommand(cleanNfsCmd)

	cleanNfsCmd.PersistentFlags().String("exports-path", "/etc/exports", "full path of file to wipe")
}
