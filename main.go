package main

import (
	"cli-todo/todo"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Load todos from file
	todos := todo.LoadTodos("data/todos.json")

	//Command line Arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo-cli <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli add <todo-title> <todo-description>")
			return
		}
		title := os.Args[2]
		description := os.Args[3]
		todos := todo.AddTodo(todos, title, description)
		todo.SaveTodos(todos, "data/todos.json")
		fmt.Println("Todo added successfully")
	case "list":
		todo.ListTodos(todos)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli delete <todo-id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID format. Please provide a number")
			return
		}
		todos = todo.DeleteTodo(todos, id)
		todo.SaveTodos(todos, "data/todos.json")
		fmt.Println("Todo deleted successfully")
	case "toggle":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli toggle <todo-id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID format. Please provide a number")
			return
		}
		todos = todo.ToggleTodoStatus(todos, id)
		todo.SaveTodos(todos, "data/todos.json")
		fmt.Println("Todo toggled successfully")
	default:
		fmt.Println("Invalid command , Available commands: add, list, delete, toggle")
	}
}
