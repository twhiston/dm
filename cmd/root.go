// Copyright © 2016 Tom Whiston <tom.whiston@gmail.com>
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
	"os"

	"bytes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os/exec"
)

var cfgFilePath string

var (
	// VERSION - The current version number, set from the main.go file
	VERSION string
	// STACK_VERSION - The version of the current stack file
	STACK_VERSION string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dm",
	Short: "Docker for mac bootstrapper",
	Long: `Docker for Mac bootstrapper
	Sets up NFS shares,
	adds local containers
	sets up socat for xdebug
	sets up loopback for phpstorm docker integration`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//
	//},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, stackVersion string) {
	VERSION = version
	STACK_VERSION = stackVersion
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "~/.dm/config.yml", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFilePath != "" {
		// enable ability to specify config file via flag
		viper.SetConfigFile(cfgFilePath)
	}

	configPath := getConfigPath()
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")
	viper.AddConfigPath(configPath) // adding home directory as first search path
	viper.AutomaticEnv()            // read in environment variables that match

	viper.Set("version", VERSION)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		viper.SetDefault("share_dir", userHomeDir())
		output := RunScript("whoami")
		viper.SetDefault("whoami", output)
		uid := RunScript("id", "-u")
		viper.SetDefault("uid", uid)
		group := RunScript("id", "-g")
		viper.SetDefault("group", group)
		viper.SetDefault("data_dir", "/Users/Shared/.dm")
		saveConfig()
	}

	// This is annoyingly a special case at the moment, really this should not be restricted to stack
	// but should be generic so any command can have an upgrade path for it's assets
	if viper.GetString("stack_version") != STACK_VERSION {

		fmt.Println("\nstack version in dm has changed, your dm stack file will be updated")
		fmt.Println("old stack version:", viper.GetString("stack_version"))
		fmt.Println("new stack version:", STACK_VERSION)
		fmt.Println()

		//make a backup of the existing stack, so user could merge back in changes later
		dir := viper.GetString("data_dir")
		backupName := dir + "/dm." + viper.GetString("stack_version") + ".yml.bck"

		fmt.Println("your current stack file will be backed up to", backupName)

		err := copyFile(dir+"/dm.yml", backupName)
		if err != nil {
			fmt.Println("Could not back up stack", err.Error())
			return
		}

		//Upgrade the stack file with the new stack
		data := GetAsset("dm.yml")
		WriteAsset(dir+"/dm.yml", data)
		viper.Set("stack_version", STACK_VERSION)
		saveConfig()
		//noinspection GoPlaceholderCount
		fmt.Println(`If dm is currently running you should run "dm stop",
copy any custom stack elements from the file backup to the new stack file
and then run "dm start" to bring up the new stack`)
		fmt.Println()
	}

}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func copyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

//RunScript runs a script by name, passing in args.
//This will either fail and exit completely or will return output
func RunScript(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + ":" + stderr.String())
		os.Exit(1)
	}
	output := out.String()
	//fmt.Print(output)
	return output
}

//TODO - REPLACE ERRORS WITH HANDLEERROR

//HandleError prints to stderror and then exists the program if not soft
func HandleError(err error, soft bool) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if !soft {
			os.Exit(1)
		}
	}
}
