package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetAsset(name string) []byte {
	data, err := Asset("assets/" + name)
	if err != nil {
		fmt.Println("Asset " + name + " missing. Please report this issue to your devops team")
		os.Exit(1)
	}
	return data
}

func WriteAsset(location string, asset []byte) {

	err := ioutil.WriteFile(location, []byte(asset), 0644)
	if err != nil {
		fmt.Println("Could not write file: " + location)
		os.Exit(1)
	}
}
