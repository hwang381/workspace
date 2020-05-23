package cmd

import (
	"flag"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var WorkspaceName string

var RootCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Multi-repo context switcher",
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&WorkspaceName, "workspace", "w", "default", "workspace name")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}
