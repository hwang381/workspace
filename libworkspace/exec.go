package libworkspace

import (
	"fmt"
	"github.com/golang/glog"
	"os"
	"os/exec"
	"sync"

	"github.com/kvz/logstreamer"
)

func executePostSwitchCommand(repoId, repoPath string, command Command) error {
	stdoutStreamer := logstreamer.NewLogstreamerForStdout(fmt.Sprintf("[%v][stdout]", repoId))
	defer func() {
		err := stdoutStreamer.Close()
		if err != nil {
			glog.Errorf("Fail to close stdout stream, %v\n", err)
		}
	}()
	stderrStreamer := logstreamer.NewLogstreamerForStderr(fmt.Sprintf("[%v][stderr]", repoId))
	defer func() {
		err := stderrStreamer.Close()
		if err != nil {
			glog.Errorf("Fail to close stderr stream, %v\n", err)
		}
	}()
	cmd := exec.Command(command.Exe[0], command.Exe[1:]...)
	var allEnvs []string
	for k, v := range command.Env {
		allEnvs = append(allEnvs, fmt.Sprintf("%v=%v", k, v))
	}
	for _, e := range os.Environ() {
		allEnvs = append(allEnvs, e)
	}
	cmd.Env = allEnvs
	cmd.Dir = repoPath
	cmd.Stdout = stdoutStreamer
	cmd.Stderr = stderrStreamer
	return cmd.Run()
}

func executePostSwitch(repo Repository) error {
	var commands []Command
	if repo.PostSwitchCommands != nil {
		commands = repo.PostSwitchCommands
	} else if repo.PostSwitch != nil && len(repo.PostSwitch) != 0 {
		commands = []Command{
			{
				Exe: repo.PostSwitch,
			},
		}
	}

	if len(commands) == 0 {
		glog.Infoln("No post switch action to be executed")
		return nil
	}

	for _, command := range commands {
		err := executePostSwitchCommand(
			repo.ID,
			repo.Path,
			command,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func executePostSwitchAsync(repo Repository, wg *sync.WaitGroup) error {
	go func() {
		err := executePostSwitch(repo)
		if err != nil {
			glog.Errorf("Fail to execute post switch, %v\n", err)
		}
		wg.Done()
	}()
	return nil
}
