package handlers

func RunCommand(cmd string, showAll bool) {
	switch cmd {
	case "add":
		HandleAdd()
	case "list":
		HandleList(showAll)
	case "update":
		HandleUpdate(showAll)
	case "delete":
		HandleDelete(showAll)
	case "help":
		HandleHelp()
	default:
		HandleGeneral()
	}
}
