package libworkspace

import (
	"fmt"
	"os"
	"os/exec"
)

func Open(repo Repository) error {
	if repo.OpenWith == "" {
		fmt.Printf("Not opening repo %s because openWith program is not specified\n", repo.ID)
		return nil
	}

	cmd := exec.Command(repo.OpenWith, repo.Path)
	cmd.Env = os.Environ()
	cmd.Dir = repo.Path

	err := cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Process.Release()
	if err != nil {
		return err
	}

	return nil
}
