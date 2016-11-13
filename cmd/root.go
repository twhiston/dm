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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"runtime"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dm",
	Short: "Docker for mac bootstrapper",
	Long: `Docker for Mac bootstrapper
	Sets up NFS shares,
	adds a mariadb container
	sets up socat for xdebug
	sets up loopback for phpstorm docker integration`,
// Uncomment the following line if your bare application
// has an action associated with it:
//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}


func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dm/dm.yml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


}

type defaultPaths struct {
	paths []string `yaml:",flow"`
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	configPath := getConfigPath()
	viper.SetConfigName("dm") // name of config file (without extension)
	viper.SetConfigType("yml")
	viper.AddConfigPath(configPath)  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	configPath = createDataDir(configPath)
	dPath := defaultPaths{
		[]string{configPath+":"+configPath+":0:0", userHomeDir()+"/Sites:"+ userHomeDir()+"/Sites:501:20"},
	}
	viper.SetDefault("nfs-paths",dPath.paths)


	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		createConfig()
	}
}

func createDataDir(configPath string) string {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Creating data dir: " + configPath)
		err = os.Mkdir(configPath, 0755)
		if err != nil {
			fmt.Println("Could not create config directory:" + configPath + " try running with sudo ")
		}
	}

	return configPath
}


func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}