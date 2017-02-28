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
	"os/exec"
)

var cfgFilePath string

var (
	// VERSION is set during build
	VERSION string
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
func Execute(version string) {
	VERSION = version
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "~/.dm/config.yml", "config file")
}

type defaultPaths struct {
	paths []string `yaml:",flow"`
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
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
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
}

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

//Print out an error and then die
//Standard error functionality
func HandleError(err error, soft bool) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if !soft {
			os.Exit(1)
		}
	}
}
