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

import "github.com/GianlucaGuarini/go-observable"

type Process interface {
	//Add any listeners that you need to the main observable.
	//See list of triggers in readme
	AddListeners(o *observable.Observable)
	//Check requirements should exit the whole command or do something to fix requirements issues
	//As its a listener it cant return anything, so your requirements handling needs to deal with all issues
	// or call os.Exit() with a sensible error message advising on a fix
	CheckRequirements(strict bool)
}
