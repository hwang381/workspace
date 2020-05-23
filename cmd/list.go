package cmd

import (
	"fmt"
	"github.com/hwang381/workspace/libworkspace"
	"github.com/spf13/cobra"
	"log"
	"sort"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "List all branches",
	// TODO: -w
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Reading configuration")
		config, err := libworkspace.ReadConfig(WorkspaceName)
		if err != nil {
			return err
		}

		log.Printf("Collecting branches for workspace %s\n", WorkspaceName)
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
	rootCmd.AddCommand(listCmd)
}
