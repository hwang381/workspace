package main

import (
	"log"
	"sync"
)

func collectBranches(repos []Repository) ([]string, error) {
	var allBranchNames []string
	for _, repo := range repos {
		branchNames, err := getBranchNames(repo)
		if err != nil {
			return nil, err
		}
		allBranchNames = append(allBranchNames, branchNames...)
	}

	dedup := make(map[string]bool)
	var branches []string
	for _, branchName := range allBranchNames {
		if !dedup[branchName] {
			dedup[branchName] = true
			branches = append(branches, branchName)
		}
	}

	log.Printf("Found branches %v", branches)
	return branches, nil
}

func switchToBranch(repos []Repository, targetBranch string, gitConfig GitConfig) error {
	levelOrder, err := getLevelOrder(repos)
	if err != nil {
		return err
	}
	log.Printf("Level ordering is %v", levelOrder)

	for _, level := range levelOrder {
		var wg sync.WaitGroup
		wg.Add(len(level))
		for _, repo := range level {
			log.Printf("Switching repo %v to %v", repo.ID, targetBranch)
			err := switchBranch(repo, targetBranch)
			if err != nil {
				return err
			}
			err = pull(repo, targetBranch, gitConfig)
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
