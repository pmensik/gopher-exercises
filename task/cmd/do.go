package cmd

import (
	"fmt"
	"strconv"

	"github.com/pmensik/gopher-exercises/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument ", arg)
			} else {
				err := db.DoTask(uint64(id))
				if err != nil {
					fmt.Printf("Task with id %d couldn't be finished.\n", id)
				}
			}
		}
		fmt.Println("Succesfully finished tasks.")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
