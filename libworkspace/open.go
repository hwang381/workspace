package libworkspace

import (
	"fmt"
	"github.com/golang/glog"
	"os"
	"os/exec"
	"time"
)

func Open(repo Repository) error {
	if repo.OpenWith == "" {
		fmt.Printf("Not opening repo %s because openWith program is not specified\n", repo.ID)
		return nil
	}

	glog.Infof("Running %s %s\n", repo.OpenWith, repo.Path)
	cmd := exec.Command(repo.OpenWith, repo.Path)
	cmd.Env = os.Environ()
	cmd.Dir = repo.Path

	err := cmd.Start()
	if err != nil {
		return err
	}

	time.Sleep(time.Second * 5)

	return nil
}
