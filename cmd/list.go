package cmd

import (
	"fmt"
	"github.com/hwang381/workspace/libworkspace"
	"github.com/spf13/cobra"
	"log"
	"sort"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all branches",
	// TODO: -w
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: dup
		log.Println("Reading config file")
		config, err := libworkspace.ReadConfig()
		if err != nil {
			return err
		}

		log.Println("Collecting branches")
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
