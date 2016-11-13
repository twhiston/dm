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
	o.On("start", func(cfgFilePath string) {
		fmt.Println("Mariadb says hello")
	})
}

func (n Mariadb) CheckRequirements(strict bool) error {
	err := rDocker{}.meetsRequirements()
	if strict {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}