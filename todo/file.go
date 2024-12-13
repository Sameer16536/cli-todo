package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Load todos from file
func LoadTodos(filename string) []Todo {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Todo{}
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []Todo{}
	}

	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		fmt.Println("Error unmarshaling data:", err)
		return []Todo{}
	}

	return todos
}

func SaveTodos(todos []Todo, filename string) {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
	}
}
