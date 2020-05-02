package libworkspace

import (
	"log"
	"os/exec"
	"strings"
)

func getBranchNames(repo Repository) ([]string, error) {
	cmd := exec.Command(
		"git",
		"for-each-ref",
		"--format=%(refname:short)",
		"refs/heads",
	)
	cmd.Dir = repo.Path
	stdoutBytes, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	stdout := string(stdoutBytes)
	return strings.Split(stdout, "\n"), nil
}

func hasBranchName(repo Repository, targetBranchName string) (bool, error) {
	branchNames, err := getBranchNames(repo)
	if err != nil {
		return false, err
	}
	for _, branchName := range branchNames {
		if branchName == targetBranchName {
			return true, nil
		}
	}
	return false, nil
}

func switchBranch(repo Repository, targetBranchName string) error {
	hasBranch, err := hasBranchName(repo, targetBranchName)
	if err != nil {
		return err
	}
	if !hasBranch {
		log.Printf("Repo %v does not have branch %v, switching to master", repo.Path, targetBranchName)
		targetBranchName = "master"
	}
	log.Printf("Switching %v to %v", repo.Path, targetBranchName)
	cmd := exec.Command(
		"git",
		"checkout",
		targetBranchName,
	)
	cmd.Dir = repo.Path
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func pull(repo Repository, branchName string) error {
	hasBranch, err := hasBranchName(repo, branchName)
	if err != nil {
		return err
	}
	if !hasBranch {
		log.Printf("Repo %v does not have branch %v, pulling master", repo.Path, branchName)
		branchName = "master"
	}
	log.Printf("Pulling %v with %v", repo.Path, branchName)
	cmd := exec.Command(
		"git",
		"pull",
	)
	cmd.Dir = repo.Path
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
