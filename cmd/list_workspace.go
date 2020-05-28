package cmd

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/hwang381/workspace/libworkspace"
	"github.com/spf13/cobra"
)

var listWorkspaceCmd = &cobra.Command{
	Use:     "list-workspace",
	Aliases: []string{"lw", "lsws"},
	Short:   "List all workspaces (aliases lw, lsws)",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		glog.Infoln("Reading configuration")
		configs, err := libworkspace.ReadConfigs()
		if err != nil {
			return err
		}

		for name := range configs {
			fmt.Println(name)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listWorkspaceCmd)
}
