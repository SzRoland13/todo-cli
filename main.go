package main

import (
	"os"

	"github.com/SzRoland13/todo-cli/handlers"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		handlers.HandleGeneral()
		return
	}

	handlers.RunCommand(args[0])
}
