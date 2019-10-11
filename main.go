package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

const cliUsage = `
Usage: workspace COMMAND ARGUMENTS
	List all issues: workspace l/ls/list
	Switch to issue: workspace s/sw/switch issue-key/"master"
`

func main() {
	// TODO: this should be controlled by -v
	// log.SetFlags(0)
	// log.SetOutput(ioutil.Discard)

	log.Println("Reading config file")
	config, err := readConfig()
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
		log.Println("Collecting JIRA issues from repositories")
		jiraIssues, err := collectIssues(config.Repositories)
		if err != nil {
			log.Fatal(err)
		}
		sort.Strings(jiraIssues)
		for _, jiraIssue := range jiraIssues {
			fmt.Println(jiraIssue)
		}
	} else if command == "s" || command == "sw" || command == "switch" {
		if len(args) != 1 {
			fmt.Print(cliUsage)
			return
		}
		targetJiraIssue := args[0]

		log.Printf("Switching all repos to %v", targetJiraIssue)
		err = switchToIssue(
			config.Repositories,
			targetJiraIssue,
			config.GitConfig,
		)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print(cliUsage)
	}
}
