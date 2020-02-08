package cmd

import (
	"fmt"
	"gophercises/task-manager/db"
	"strings"

	"github.com/spf13/cobra"
)

var addcmd = &cobra.Command{
	Use:   "add",
	Short: "Adds tasks to your list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your  task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addcmd)
}
