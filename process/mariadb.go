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
)

type Mariadb struct {

}

func (n Mariadb) AddListeners(o *observable.Observable) {
	o.On("start", Mariadb.start)
}

func (n Mariadb) start(cfgFilePath string) {
	fmt.Println("	---> Setup local MariaDB data folder")
	var mariaDir = cfgFilePath + "/maria"
	if _, err := os.Stat(mariaDir); os.IsNotExist(err) {
		//Install stuff if it doesnt exist
		fmt.Print("		---> Creating mariadb dir: ")
		fmt.Println(mariaDir)
		handleError(os.Mkdir(mariaDir, 0755))

		fmt.Print("		---> Creating mariadb data dir: ")
		var mariaDataDir = mariaDir + "/data"
		fmt.Println(mariaDataDir)
		handleError(os.Mkdir(mariaDataDir, 0755))

		fmt.Println("		---> Copying docker assets: ")
		data := getAsset("database/.env")
		writeAsset(mariaDir+"/.env", data)

		data = getAsset("database/maria.yml")
		writeAsset(mariaDir+"/maria.yml", data)

	}
	//docker compose up the maria db
	fmt.Println("	---> Starting mariadb container")
	RunScript("/bin/sh", "-c", "docker-compose -f "+mariaDir+"/maria.yml up -d")

	fmt.Println("	     Started mariadb, you can use it in your local environment docker-compose.yml")
	fmt.Println("          	external_links:")
	fmt.Println("              - mariadb_local")
}

func (n Mariadb) CheckRequirements(strict bool)  {
	err := rDocker{}.meetsRequirements()
	if strict {
		handleError(err)
	}

}