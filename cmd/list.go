package cmd

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/hwang381/workspace/libworkspace"
	"github.com/spf13/cobra"
	"sort"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "List all branches (aliases l, ls)",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		glog.Infoln("Reading configuration")
		config, err := libworkspace.ReadConfig(WorkspaceName)
		if err != nil {
			return err
		}

		glog.Infoln("Collecting branches for workspace %v", WorkspaceName)
		branches, err := libworkspace.CollectBranches(config.Repositories)
		if err != nil {
			return err
		}
		sort.Strings(branches)
		for _, branch := range branches {
			fmt.Println(branch)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
