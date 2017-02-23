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

package cmd

import (
	"github.com/GianlucaGuarini/go-observable"
	"github.com/twhiston/dm/process"
)

//Run this in your command to create an observer and pass it to the processes to add their listeners
//The process will add ALL its listeners, so you can have multiple triggers in your command
func SetUpListeners() *observable.Observable {

	var processes = [...]process.Process{process.Nfs{}, process.Mariadb{}}

	o := observable.New()

	for _, p := range processes {
		p.AddListeners(o)
	}
	return o
}

