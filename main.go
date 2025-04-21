package main

import (
	"os"
	"strconv"

	"github.com/SzRoland13/todo-cli/handlers"
)

func main() {
	args := os.Args[1:]

	val, _ := strconv.ParseBool(args[1])

	handlers.RunCommand(args[0], val)
}
