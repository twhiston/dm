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

	"bytes"
	"github.com/spf13/cobra"
	"os/exec"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List running docker containers",
	Long: `Exactly analogous to docker ps,
normally just called by other commands to show containers during operations`,
	Run: func(cmd *cobra.Command, args []string) {
		listContainers()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func listContainers() {
	fmt.Println("---> Listing Running Docker Containers")
	cmd := exec.Command("docker", "ps")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	fmt.Print(out.String())
}
