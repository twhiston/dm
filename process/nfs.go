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

package process

import (
	"github.com/GianlucaGuarini/go-observable"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/libgit2/git2go"
)

type Nfs struct {

}

func (n Nfs) AddListeners(o *observable.Observable) {
	o.On("start", Nfs.start)
}

func (n Nfs) start(cfgFilePath string) {
	fmt.Println("	---> Setup NFS shares")
	var nfsDir = cfgFilePath + "/d4m-nfs"

	if _, err := os.Stat(nfsDir); os.IsNotExist(err) {
		//If the directory doesn't exist then make it and clone the helper repo we are using
		fmt.Print("Creating nfs mount script dir: ")
		fmt.Println(nfsDir)
		handleError(os.Mkdir(nfsDir, 0755))
		_, err = git.Clone("https://github.com/IFSight/d4m-nfs", nfsDir, &git.CloneOptions{})
		handleError(err)
		//Now the repo is cloned copy in our unique assets
		//Get the data from the config file
		//Turn it into a text file and write it to the /etc/folder
		//data := getAsset("nfs/d4m-nfs-mounts.txt")
		//writeAsset(nfsDir+"/etc/d4m-nfs-mounts.txt", data)
	}

	//Run the command
	RunScript(nfsDir + "/d4m-nfs.sh")
}

func (n Nfs) CheckRequirements(strict bool) {
	err := rOsx{}.meetsRequirements()
	if strict {
		handleError(err)
	}

	k := viper.AllKeys()
	fmt.Println(k)

}