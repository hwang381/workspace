package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var WorkspaceName string

var rootCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Multi-repo context switcher",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&WorkspaceName, "workspace", "w", "default", "workspace name")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
