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
	"fmt"
	"io/ioutil"
	"os"
)

func getAsset(name string) []byte {
	data, err := Asset("data/" + name)
	if err != nil {
		fmt.Println("Asset " + name + " missing. Please report this issue to your devops team")
		os.Exit(1)
	}
	return data
}

func writeAsset(location string, asset []byte) {

	err := ioutil.WriteFile(location, []byte(asset), 0644)
	if err != nil {
		fmt.Println("Could not write mariadb env file")
		os.Exit(1)
	}
}
