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
