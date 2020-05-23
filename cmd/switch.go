package cmd

import (
	"github.com/hwang381/workspace/libworkspace"
	"github.com/spf13/cobra"
	"log"
)

var switchCmd = &cobra.Command{
	Use:     "switch",
	Aliases: []string{"s", "sw"},
	Short:   "List all branches",
	// TODO: -w
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: dup
		log.Println("Reading config file")
		configs, err := libworkspace.ReadConfigs()
		if err != nil {
			return err
		}

		targetBranch := args[0]
		log.Printf("Switching all repos to %v", targetBranch)
		config := configs["default"]
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
