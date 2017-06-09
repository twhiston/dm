package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
)

//GetAsset takes a named asset and loads it or fails and exists
func GetAsset(name string) []byte {
	data, err := Asset("assets/" + name)
	if err != nil {
		fmt.Println("Asset " + name + " missing. Please report this issue to your devops team")
		os.Exit(1)
	}
	return data
}

//WriteAsset writes a byte array to a location
func WriteAsset(location string, asset []byte) {

	err := ioutil.WriteFile(location, asset, 0644)
	if err != nil {
		fmt.Println("Could not write file: " + location)
		os.Exit(1)
	}
}
