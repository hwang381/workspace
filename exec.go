package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/kvz/logstreamer"
)

func executePostSwitch(repo Repository) error {
	if repo.PostSwitch == nil || len(repo.PostSwitch) == 0 {
		log.Printf("no post switch action to be executed")
		return nil
	}
	stdoutStreamer := logstreamer.NewLogstreamerForStdout(fmt.Sprintf("[%v][stdout]", repo.ID))
	defer func() {
		err := stdoutStreamer.Close()
		if err != nil {
			log.Printf("fail to close stdout stream, %v", err)
		}
	}()
	stderrStreamer := logstreamer.NewLogstreamerForStderr(fmt.Sprintf("[%v][stderr]", repo.ID))
	defer func() {
		err := stderrStreamer.Close()
		if err != nil {
			log.Printf("fail to close stderr stream, %v", err)
		}
	}()
	cmd := exec.Command(repo.PostSwitch[0], repo.PostSwitch[1:]...)
	cmd.Env = os.Environ()
	cmd.Dir = repo.Path
	cmd.Stdout = stdoutStreamer
	cmd.Stderr = stderrStreamer
	return cmd.Run()
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
