package cmd

import (
	"fmt"
	"gophercises/task-manager/db"
	"strconv"

	"github.com/spf13/cobra"
)

var docmd = &cobra.Command{
	Use:   "do",
	Short: "Mark tasks as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse arguments")
			} else {
				ids = append(ids, id)
			}
			tasks, err := db.ViewTasks()
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			for _, id := range ids {
				if id < 0 || id > len(tasks) {
					fmt.Println("Invalid Task number:", id)
					continue
				}
				task := tasks[id-1]
				err := db.DeleteTask(task.Key)
				if err != nil {
					fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err.Error())
				} else{
					fmt.Printf("Failed to mark \"%d\" as completed.", id)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(docmd)
}
