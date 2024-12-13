# Todo CLI Application

A simple Command-Line Interface (CLI) Todo application built in Go, allowing you to manage your tasks effectively with options to add, delete, toggle, and display tasks in a tabular format.

---

## **Features**

- Add new todos.
- List all todos in a tabular format with status.
- Toggle the completion status of a todo.
- Delete a todo.
- Persistent storage using a JSON file.

---

## **Technologies Used**

- **Go**: The programming language used to build the application.
- **Tablewriter**: A library for creating beautiful ASCII tables.
- **Fatih/Color**: A library for adding colors to terminal output.

---

## **Folder Structure**

```plaintext
todo-cli/
├── main.go              # Entry point of the application
├── todo/
│   ├── todo.go          # Todo struct and core functions (CRUD)
│   ├── file.go          # File I/O functions for persistence
│   └── display.go       # Functions for displaying todos
├── data/
│   └── todos.json       # Data file for storing todos
├── go.mod               # Go module file
└── go.sum               # Go dependencies file
```

---

## **Getting Started**

### **Prerequisites**

- Install Go: [https://golang.org/dl/](https://golang.org/dl/)

---

### **Setup**

1. Clone the repository or create a project folder:
   ```bash
   mkdir todo-cli && cd todo-cli
   ```

2. Initialize a Go module:
   ```bash
   go mod init todo-cli
   ```

3. Install dependencies:
   ```bash
   go get github.com/olekukonko/tablewriter
   go get github.com/fatih/color
   ```

4. Create the folder structure as shown above.

5. Initialize an empty JSON file for data:
   ```bash
   mkdir data
   touch data/todos.json
   ```
   Add the following content to `todos.json`:
   ```json
   []
   ```

---

## **Commands**

### **Add a Todo**
```bash
./todo-cli add "<todo-title>"
```
Example:
```bash
./todo-cli add "Learn Go"
```

### **List Todos**
```bash
./todo-cli list
```

### **Toggle Todo Status**
```bash
./todo-cli toggle <index>
```
Example:
```bash
./todo-cli toggle 1
```

### **Delete a Todo**
```bash
./todo-cli delete <index>
```
Example:
```bash
./todo-cli delete 2
```

---

## **Run the Application**

To run the application directly:
```bash
go run main.go <command> [arguments]
```

Example:
```bash
go run main.go add "Learn Go"
```

---

## **Build the Executable**

To build the application into an executable:
```bash
go build -o todo-cli
```

Run the executable:
```bash
./todo-cli <command> [arguments]
```

Example:
```bash
./todo-cli list
```

---

## **Example Usage**

1. Add Todos:
   ```bash
   ./todo-cli add "Learn Go"
   ./todo-cli add "Build CLI app"
   ```

2. List Todos:
   ```bash
   ./todo-cli list
   ```
   Output:
   ```plaintext
   +-------+----------------+-------------+
   | Index | Title          | Status      |
   +-------+----------------+-------------+
   | 1     | Learn Go       | ❌ Incomplete |
   | 2     | Build CLI app  | ❌ Incomplete |
   +-------+----------------+-------------+
   ```

3. Toggle a Todo's Status:
   ```bash
   ./todo-cli toggle 1
   ```

4. Delete a Todo:
   ```bash
   ./todo-cli delete 2
   ```

---

## **License**

This project is open-source and available under the MIT License.

