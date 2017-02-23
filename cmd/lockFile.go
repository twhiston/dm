package cmd

import (
	"os"
	"io/ioutil"
	"fmt"
)

func getLockFileAbsolutePath(cfgFilePath string) string {
	return cfgFilePath + "/.lock"
}

func createLockFile(cfgFilePath string, forceFlag bool) {
	lockFile := getLockFileAbsolutePath(cfgFilePath)

	if _, err := os.Stat(lockFile); os.IsNotExist(err) || forceFlag == true {

		err := ioutil.WriteFile(lockFile, []byte("lock"), 0644)
		if err != nil {
			fmt.Println("Could not write lock file")
			os.Exit(1)
		}
	} else {
		fmt.Println("	Cannot start, lock file already exists, run -stop first")
		os.Exit(1)
	}
}

func deleteLockFile(cfgFilePath string) {
	lockFile := getLockFileAbsolutePath(cfgFilePath)
	_, err := os.Stat(lockFile)
	if err == nil {
		os.Remove(lockFile)
	}
}
