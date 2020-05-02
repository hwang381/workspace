package main

import (
	"fmt"
	"github.com/hwang381/workspace/libworkspace"
	"log"
	"os"
	"os/exec"
	"sort"
)

const cliUsage = `
Usage: workspace COMMAND ARGUMENTS
	List all branches: workspace l/ls/list
	Switch to branch: workspace s/sw/switch branch-name/"master"
`

func main() {
	// TODO: this should be controlled by -v
	// log.SetFlags(0)
	// log.SetOutput(ioutil.Discard)
	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Reading config file")
	config, err := libworkspace.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cliArgs := os.Args[1:]
	if len(cliArgs) < 1 {
		fmt.Print(cliUsage)
		return
	}
	command := cliArgs[0]
	args := cliArgs[1:]

	if command == "l" || command == "ls" || command == "list" {
		log.Println("Collecting branches from repositories")
		branches, err := libworkspace.CollectBranches(config.Repositories)
		if err != nil {
			log.Fatal(err)
		}
		sort.Strings(branches)
		for _, branch := range branches {
			fmt.Println(branch)
		}
	} else if command == "s" || command == "sw" || command == "switch" {
		if len(args) != 1 {
			fmt.Print(cliUsage)
			return
		}
		targetBranch := args[0]

		log.Printf("Switching all repos to %v", targetBranch)
		err = libworkspace.SwitchToBranch(
			config.Repositories,
			targetBranch,
		)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print(cliUsage)
	}
}
