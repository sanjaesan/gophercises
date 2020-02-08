package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Tasker is a task manager",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
