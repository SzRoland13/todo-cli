package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/SzRoland13/todo-cli/todo"
)

func HandleAdd() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Description: ")
	desc, _ := reader.ReadString('\n')

	fmt.Print("Priority (1=LOW, 2=MEDIUM, 3=HIGH): ")
	var p int
	fmt.Scanf("%d\n", &p)

	fmt.Print("Due Date (YYYY-MM-DD): ")
	var dueStr string
	fmt.Scanf("%s\n", &dueStr)
	due, _ := time.Parse("2006-01-02", dueStr)

	newTodo := todo.Todo{
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(desc),
		Priority:    todo.Priority(p),
		Progress:    todo.TO_DO,
		DueDate:     due,
	}

	todo.AddTodo(newTodo)

	fmt.Println("Todo added!")
}

func HandleGeneral() {
	fmt.Println("Choose a command: add / list / update / delete / help")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")

	cmd, _ := reader.ReadString('\n')

	cmd = strings.TrimSpace(cmd)

	RunCommand(cmd, false)

}

func HandleList(showAll bool) {
	if showAll {
		todo.ListAllTodos()
	} else {
		todo.ListActiveTodos()
	}
}

func HandleUpdate(showAll bool) {
	fmt.Println("Choose which Todo to update")

	HandleList(showAll)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("ID: ")

	idStr, _ := reader.ReadString('\n')

	idStr = strings.TrimSpace(idStr)

	ID, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil || !isIdInList(uint(ID), showAll) {
		fmt.Println("Invalid ID was passed! Try again with a valid ID.")
		return
	}

	t := todo.GetTodoByID(uint(ID))

	if t == nil {
		fmt.Println("Todo not found")
		return
	}

	fmt.Println("Which field do you want to update? (Title, Description, Priority, Progress, DueDate)")
	fmt.Print("> ")

	field, _ := reader.ReadString('\n')

	field = strings.TrimSpace(strings.ToLower(field))

	switch field {
	case "title":
		fmt.Print("New Title: ")
		newVal, _ := reader.ReadString('\n')
		t.Title = strings.TrimSpace(newVal)

	case "description":
		fmt.Print("New Description: ")
		newVal, _ := reader.ReadString('\n')
		t.Description = strings.TrimSpace(newVal)

	case "priority":
		fmt.Print("New Priority (1=LOW, 2=MEDIUM, 3=HIGH): ")
		newVal, _ := reader.ReadString('\n')
		newVal = strings.TrimSpace(newVal)
		num, err := strconv.Atoi(newVal)
		if err != nil || num < 1 || num > 3 {
			fmt.Println("Invalid priority!")
			return
		}
		t.Priority = todo.Priority(num)

	case "progress":
		fmt.Print("New Progress (1=TO_DO, 2=IN_PROGRESS, 3=ON_HOLD, 4=COMPLETED): ")
		newVal, _ := reader.ReadString('\n')
		newVal = strings.TrimSpace(newVal)
		num, err := strconv.Atoi(newVal)
		if err != nil || num < 1 || num > 4 {
			fmt.Println("Invalid progress!")
			return
		}
		t.Progress = todo.Progress(num)

	case "duedate":
		fmt.Print("New Due Date (YYYY-MM-DD): ")
		newVal, _ := reader.ReadString('\n')
		newVal = strings.TrimSpace(newVal)
		due, err := time.Parse("2006-01-02", newVal)
		if err != nil {
			fmt.Println("Invalid date format!")
			return
		}
		t.DueDate = due

	default:
		fmt.Println("Unknown field.")
		return
	}

	_ = todo.SaveTodos()
	fmt.Println("Todo updated successfully.")

}

func isIdInList(ID uint, showAll bool) bool {
	for _, element := range todo.GetTodos() {
		if element.ID == ID {
			return true
		}
	}

	return false
}

func HandleHelp() {
	fmt.Print(`
Todo CLI - Commands:

  add         Add a new todo interactively
  list        List all todos
  help        Show this help message

Example:
  go run main.go add
  go run main.go list
`)
}

func HandleDelete(showAll bool) {
	fmt.Println("Choose which Todo to delete")
	HandleList(showAll)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("ID: ")

	idStr, _ := reader.ReadString('\n')

	idStr = strings.TrimSpace(idStr)

	ID, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil || !isIdInList(uint(ID), showAll) {
		fmt.Println("Invalid ID was passed! Try again with a valid ID.")
		return
	}

	t := todo.GetTodoByID(uint(ID))

	if t == nil {
		fmt.Println("Todo not found")
		return
	}

	t.IsDeleted = true

	_ = todo.SaveTodos()
	fmt.Println("Todo soft deleted successfully.")
}
