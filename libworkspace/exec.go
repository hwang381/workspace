package libworkspace

import (
	"fmt"
	"log"
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
			log.Printf("fail to close stdout stream, %v", err)
		}
	}()
	stderrStreamer := logstreamer.NewLogstreamerForStderr(fmt.Sprintf("[%v][stderr]", repoId))
	defer func() {
		err := stderrStreamer.Close()
		if err != nil {
			log.Printf("fail to close stderr stream, %v", err)
		}
	}()
	cmd := exec.Command(command.Exe[0], command.Exe[1:]...)
	cmd.Env = os.Environ()
	if command.Environ != nil {
		for key, value := range command.Environ {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
		}
	}
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
		log.Printf("no post switch action to be executed")
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
			log.Panicf("fail to execute post switch, %v", err)
		}
		wg.Done()
	}()
	return nil
}
