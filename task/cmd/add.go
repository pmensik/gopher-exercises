package cmd

import (
	"fmt"
	"strings"

	"github.com/pmensik/gopher-exercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if err := db.AddTask(task); err != nil {
			fmt.Println("Couldn't add a task")
		} else {
			fmt.Printf("Added \"%s\" to your list\n", task)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
