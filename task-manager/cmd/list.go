package cmd

import (
	"fmt"
	"gophercises/task-manager/db"
	"os"

	"github.com/spf13/cobra"
)

var listcmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks on your list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ViewTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete")
			return
		}
		fmt.Println("You have the following tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s, key=%d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(listcmd)
}
