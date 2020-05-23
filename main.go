package main

import (
	"github.com/hwang381/workspace/cmd"
	"log"
	"os/exec"
)

func main() {
	// TODO: this should be controlled by -v
	// log.SetFlags(0)
	// log.SetOutput(ioutil.Discard)
	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
