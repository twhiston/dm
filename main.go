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

package main

import (
	"github.com/twhiston/dm/cmd"
)

var (
	//VERSION - Current application version
	VERSION = "1.0.0-beta3"
	//STACK_VERSION - Current stack file version
	STACK_VERSION = "1.1"
)

func main() {
	cmd.Execute(VERSION, STACK_VERSION)
}
