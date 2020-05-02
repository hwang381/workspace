package main

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

const branchRefPrefix = "refs/heads/"

func getAuthentication(gitConfig GitConfig) (*gitssh.PublicKeys, error) {
	log.Printf("Reading private key from %v", gitConfig.PrivateKeyPath)
	auth, err := gitssh.NewPublicKeysFromFile(
		"git",
		gitConfig.PrivateKeyPath,
		gitConfig.KeyPairPassphrase,
	)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func getBranchNames(repo Repository) ([]string, error) {
	gitRepo, err := git.PlainOpen(repo.Path)
	if err != nil {
		return nil, err
	}
	branchesIter, err := gitRepo.Branches()
	if err != nil {
		return nil, err
	}
	var branchNames []string
	err = branchesIter.ForEach(func(ref *plumbing.Reference) error {
		branchRef := string(ref.Name())
		branchName := branchRef[len(branchRefPrefix):]
		branchNames = append(branchNames, branchName)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return branchNames, nil
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
	gitRepo, err := git.PlainOpen(repo.Path)
	if err != nil {
		return err
	}
	workTree, err := gitRepo.Worktree()
	if err != nil {
		return err
	}
	log.Printf("Switching %v to %v", repo.Path, targetBranchName)
	err = workTree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branchRefPrefix + targetBranchName),
		Force:  false,
	})
	if err != nil {
		return err
	}
	return nil
}

func pull(repo Repository, branchName string, gitConfig GitConfig) error {
	hasBranch, err := hasBranchName(repo, branchName)
	if err != nil {
		return err
	}
	if !hasBranch {
		log.Printf("Repo %v does not have branch %v, pulling master", repo.Path, branchName)
		branchName = "master"
	}
	gitRepo, err := git.PlainOpen(repo.Path)
	if err != nil {
		return err
	}
	workTree, err := gitRepo.Worktree()
	if err != nil {
		return err
	}
	auth, err := getAuthentication(gitConfig)
	if err != nil {
		return err
	}
	err = workTree.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(branchRefPrefix + branchName),
		SingleBranch:  true,
		Auth:          auth,
		Force:         false,
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}
