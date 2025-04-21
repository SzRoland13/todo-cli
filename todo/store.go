package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

const filePath = "data/todos.json"

var todos []Todo

func LoadTodos() error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			todos = []Todo{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &todos)
}

func SaveTodos() error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func AddTodo(t Todo) {
	var maxID uint
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	t.ID = maxID + 1

	todos = append(todos, t)
}

func ListActiveTodos() {
	for _, t := range todos {
		if !t.IsDeleted || t.Progress != 4 {
			fmt.Printf("%d. %s: %s [%s] (%s) 📅 %s\n", t.ID, t.Title, t.Description, t.Priority.String(), t.Progress.String(), t.DueDate.Format("2006-01-02"))
		}
	}
}

func ListAllTodos() {
	for _, t := range todos {
		fmt.Printf("%d. %s: %s [%s] (%s) 📅 %s\n", t.ID, t.Title, t.Description, t.Priority.String(), t.Progress.String(), t.DueDate.Format("2006-01-02"))
	}
}

func GetTodos() []Todo {
	return todos
}

func GetTodoByID(ID uint) *Todo {
	for i, val := range todos {
		if val.ID == ID && !val.IsDeleted {
			return &todos[i]
		}
	}
	return nil
}
