package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/hwang381/workspace/cmd"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()

	_, err := exec.LookPath("git")
	if err != nil {
		glog.Errorln("git cannot be found, %v", err)
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		glog.Errorln(err)
		os.Exit(1)
	}
}
