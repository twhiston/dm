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
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"time"
)

func getLockFileAbsolutePath() string {
	cp := viper.GetString("lock_path")
	if cp == "" {
		cp = "/tmp"
	}
	return cp + "/.dmlock"
}

func createLockFile(forceFlag bool) {
	lockFile := getLockFileAbsolutePath()

	if _, err := os.Stat(lockFile); os.IsNotExist(err) || forceFlag {

		err := ioutil.WriteFile(lockFile, makeTimestamp(), 0644)
		if err != nil {
			fmt.Println("Could not write lock file: " + lockFile)
			os.Exit(1)
		}
	} else {
		fmt.Println("Cannot start, lock file already exists, run stop first or use -f --force flag")
		os.Exit(1)
	}
}

func makeTimestamp() []byte {
	stamp := time.Now().Unix()
	return []byte(fmt.Sprintf("%d", stamp))
}

func deleteLockFile() {
	lockFile := getLockFileAbsolutePath()
	_, err := os.Stat(lockFile)
	if err == nil {
		os.Remove(lockFile)
	}
}
