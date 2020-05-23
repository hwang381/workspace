package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Multi-repo context switcher",
	Long: "workspace manages your multi-repo project and makes it easy to context switch them for different ongoing " +
		"work items while respecting inter-repo dependencies",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
