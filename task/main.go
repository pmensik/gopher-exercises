package main

import (
	"log"

	"github.com/pmensik/gopher-exercises/task/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
