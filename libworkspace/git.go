package libworkspace

import (
	"github.com/golang/glog"
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
		glog.Infof("Repo %s does not have branch %s, switching to master\n", repo.Path, targetBranchName)
		targetBranchName = "master"
	}
	glog.Infof("Switching %s to %s\n", repo.Path, targetBranchName)
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
		glog.Infof("Repo %s does not have branch %s, pulling master\n", repo.Path, branchName)
		branchName = "master"
	}
	glog.Infof("Pulling %s with %s\n", repo.Path, branchName)
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
