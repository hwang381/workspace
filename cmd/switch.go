package cmd

import (
	"github.com/golang/glog"
	"github.com/hwang381/workspace/libworkspace"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:     "switch",
	Aliases: []string{"s", "sw"},
	Short:   "Switch all repositories into a branch (aliases s, sw)",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		glog.Infoln("Reading configuration")
		config, err := libworkspace.ReadConfig(WorkspaceName)
		if err != nil {
			return err
		}

		targetBranch := args[0]
		glog.Infof("Switching all repos to %s for workspace %s\n", targetBranch, WorkspaceName)
		err = libworkspace.SwitchToBranch(
			config.Repositories,
			targetBranch,
		)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(switchCmd)
}
