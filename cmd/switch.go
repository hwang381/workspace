package cmd

import (
	"github.com/hwang381/workspace/libworkspace"
	"github.com/spf13/cobra"
	"log"
)

var switchCmd = &cobra.Command{
	Use:     "switch",
	Aliases: []string{"s", "sw"},
	Short:   "Switch branch of all repositories into a branch",
	// TODO: -w
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Reading configurations")
		configs, err := libworkspace.ReadConfigs()
		if err != nil {
			return err
		}

		targetBranch := args[0]
		log.Printf("Switching all repos to %s for workspace %s\n", targetBranch, WorkspaceName)
		config := configs[WorkspaceName]
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
	rootCmd.AddCommand(switchCmd)
}
