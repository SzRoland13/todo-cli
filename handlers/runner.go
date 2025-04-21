package handlers

import "fmt"

func RunCommand(cmd string) {
	switch cmd {
	case "add":
		HandleAdd()
	case "list":
		HandleList()
	case "update":
		HandleUpdate()
	case "delete":
		HandleDelete()
	case "help":
		HandleHelp()
	default:
		fmt.Println("Unkown command. Type `help` for usage")
	}
}
