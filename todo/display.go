package todo

import (
	"fmt"

	"os"

	"github.com/olekukonko/tablewriter"
)

func ListTodos(todos []Todo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Description", "Status"})

	for i, todo := range todos {
		status := "❌ Incomplete"
		if todo.Completed {
			status = "✅ Completed"
		}
		table.Append([]string{fmt.Sprintf("%d", i+1), todo.Title, todo.Description, status})
	}

	table.SetBorder(true)
	table.SetRowLine(true)
	table.Render()
}
