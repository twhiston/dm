package cmd

import (
	"github.com/spf13/viper"
	"testing"
)

func TestLockDir(t *testing.T) {

	defaultPath := "/tmp/.dmlock"

	path := getLockFileAbsolutePath()
	if path != defaultPath {
		t.Fatal("without viper config path should be", defaultPath)
	}

	newPath := "/testpath"
	viper.Set("lock_path", newPath)

	path = getLockFileAbsolutePath()
	if path != newPath+"/.dmlock" {
		t.Fatal("without viper config path should be", newPath)
	}

}
