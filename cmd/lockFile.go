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
