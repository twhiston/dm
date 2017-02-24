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
	"os/exec"
	"bytes"
	"fmt"
	"os"
)

func RunScript(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + ":" + stderr.String())
		os.Exit(1)
	}
	output := out.String()
	//fmt.Print(output)
	return output
}

//Print out an error and then die
//Standard error functionality
func HandleError(err error) {
	if(err != nil) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}