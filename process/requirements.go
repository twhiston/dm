package process

import (
	"errors"
	"github.com/matishsiao/goInfo"
	"os"
	"os/exec"
)

type Requirement interface {
	meetsRequirements() error
}

type rOsx struct{}

func (r rOsx) meetsRequirements() error {
	gi := goInfo.GetInfo()
	if gi.GoOS != "darwin" {
		return errors.New("This Installer cannot be run without OSX")
	}
	return nil
}

type rDocker struct{}

func (r rDocker) meetsRequirements() error {
	if err := exec.Command("docker").Run(); err != nil {
		return errors.New("	---> Could not find docker on your system\n	     Please install Docker for Mac before running this program\n 	     https://docs.docker.com/docker-for-mac")
	}
	_, exists := os.LookupEnv("DOCKER_HOST")
	if exists {
		return errors.New("	Found something checking for docker envs.\n	This suggests you have the old docker toolbox, please install docker for mac and unset docker vars\n 	See: https://docs.docker.com/docker-for-mac/docker-toolbox/#/setting-up-to-run-docker-for-mac")
	}
	return nil
}

type rSocat struct{}

func (r rSocat) meetsRequirements() error {
	if err := exec.Command("socat", "-V").Run(); err != nil {
		return errors.New("	---> Could not find socat on your system\n		Try `brew install socat`")
	}
	return nil
}
