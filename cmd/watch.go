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
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type WatchList []Watcher

type Watcher struct {
	Pattern   string   `json:"pattern"`
	Commands  []string `json:"commands"`
	Recursive bool     `json:"recursive"`
}

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch a folder, or set of folders and do something",
	Long: `Expects to find a file of directories to watch and steps to execute watchers.yml
Can also be specified with the --watchfile option
Uses https://github.com/fsnotify/fsnotify under the hood for watching, so see the github page for information if issues occur`,
	Run: func(cmd *cobra.Command, args []string) {

		loc, _ := cmd.PersistentFlags().GetString("location")
		rcmd, _ := cmd.PersistentFlags().GetString("command")

		if (loc == "" && rcmd != "") || (loc != "" && rcmd == "") {
			//incorrect input
			fmt.Println("If you set location you must also set command")
			return
		}

		if loc != "" && rcmd != "" {
			//In this case we just do a single watch
			doWatch(loc, rcmd)
			return
		}

		//if watchfile is set run with that

		//if not look in current folder for file and run with that

		//if still nothing return

	},
}

// Needs to ignore ___jb_tmp___ files
// Would be cool to have a real ignore functionality
func doWatch(loc string, cmd string) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				//Editing a file triggers a create event
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("DM Watcher Triggered")
					fmt.Println("event:", event)
					output, err := exec.Command(cmd).Output()
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Printf("%s", output)
					}
					fmt.Println()
				}

			case err := <-watcher.Errors:
				log.Println("error:", err)
				fmt.Println()
			}
		}
	}()

	//err = watcher.Add(loc)
	watchRecursive(loc, false, watcher)
	if err != nil {
		log.Fatal(err)
	}
	<-done

}

// watchRecursive adds all directories under the given one to the watch list.
// this is probably a very racey process. What if a file is added to a folder before we get the watch added?
func watchRecursive(path string, unWatch bool, watcher *fsnotify.Watcher) error {
	err := filepath.Walk(path, func(walkPath string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			if unWatch {
				if err = watcher.Remove(walkPath); err != nil {
					return err
				}
			} else {
				if err = watcher.Add(walkPath); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func init() {
	RootCmd.AddCommand(watchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	watchCmd.PersistentFlags().String("watchfile", "", "An optional watchfile including path")
	watchCmd.PersistentFlags().String("location", "", "Location is a pattern to watch, will override a local or specified watchfile. If set command must also be specified")
	watchCmd.PersistentFlags().String("command", "", "A single command to run when changes are detected in location, will override a local or specified watchfile. If set location must also be specified")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
