package cmd

import (
	"fmt"

	"github.com/pmensik/gopher-exercises/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Couldn't list tasks")
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks in the list.")
		} else {
			for _, t := range tasks {
				fmt.Printf("%d : %s\n", t.Id, t.Text)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
