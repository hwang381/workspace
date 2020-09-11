package cmd

import (
	"github.com/golang/glog"
	"github.com/hwang381/workspace/libworkspace"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o", "op"},
	Short:   "Open a workspace with the specified programs",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		glog.Infoln("Reading configuration")
		config, err := libworkspace.ReadConfig(WorkspaceName)
		if err != nil {
			return err
		}

		glog.Infoln("Opening %s\n", WorkspaceName)
		for _, repo := range config.Repositories {
			err = libworkspace.Open(repo)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(openCmd)
}
