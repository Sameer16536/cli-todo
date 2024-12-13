package todo

import "fmt"

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

// Add new Todo
func AddTodo(todos []Todo, title string, description string) []Todo {
	newTodo := Todo{
		ID:          len(todos) + 1,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	return append(todos, newTodo)
}

// Delete a todo
func DeleteTodo(todos []Todo, id int) []Todo {
	if id < 1 || id > len(todos) {
		fmt.Println("Invalid todo ID")
		return todos
	}
	return append(todos[:id-1], todos[id:]...)
}

//Toggle todo status
func ToggleTodoStatus(todos []Todo, id int) []Todo {
	if id < 1 || id > len(todos) {
		fmt.Println("Invalid todo ID")
		return todos
	}
	todos[id-1].Completed = !todos[id-1].Completed
	return todos
}
