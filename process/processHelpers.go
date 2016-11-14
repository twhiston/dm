package process

import (
	"os/exec"
	"bytes"
	"fmt"
	"os"
)

func runScript(name string, args ...string) {
	cmd := exec.Command(name, args...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + ":" + stderr.String())
		os.Exit(1)
	}
	fmt.Print(out.String())
}

//Print out an error and then die
//Standard error functionality
func handleError(err error) {
	if(err != nil) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}