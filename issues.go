package main

import (
	"log"
	"strings"
	"sync"
)

const branchNameJiraPrefix = "jira/"

func collectIssues(repos []Repository) ([]string, error) {
	var allBranchNames []string
	for _, repo := range repos {
		branchNames, err := getBranchNames(repo)
		if err != nil {
			return nil, err
		}
		allBranchNames = append(allBranchNames, branchNames...)
	}

	dedup := make(map[string]bool)
	var jiraIssues []string
	for _, branchName := range allBranchNames {
		if strings.HasPrefix(branchName, branchNameJiraPrefix) {
			jiraIssue := branchName[len(branchNameJiraPrefix):]
			if !dedup[jiraIssue] {
				dedup[jiraIssue] = true
				jiraIssues = append(jiraIssues, jiraIssue)
			}
		}
	}

	log.Printf("Found issues %v", jiraIssues)
	return jiraIssues, nil
}

func switchToIssue(repos []Repository, targetIssue string, gitConfig GitConfig) error {
	levelOrder, err := getLevelOrder(repos)
	if err != nil {
		return err
	}
	log.Printf("Level ordering is %v", levelOrder)

	for _, level := range levelOrder {
		var wg sync.WaitGroup
		wg.Add(len(level))
		for _, repo := range level {
			log.Printf("Switching repo %v to %v", repo.ID, targetIssue)
			var targetBranchName string
			if targetIssue == "master" {
				targetBranchName = "master"
			} else {
				targetBranchName = branchNameJiraPrefix + targetIssue
			}
			err := switchBranch(repo, targetBranchName)
			if err != nil {
				return err
			}
			err = pull(repo, targetBranchName, gitConfig)
			if err != nil {
				return err
			}

			log.Printf("Executing post switch for repo %v", repo.ID)
			err = executePostSwitchAsync(repo, &wg)
			if err != nil {
				return err
			}
		}
		wg.Wait()
	}

	return nil
}
