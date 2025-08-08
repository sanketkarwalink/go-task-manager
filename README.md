# Go CLI Task Manager 📝

A simple command-line task manager built with Go. You can add, list, complete, and delete tasks — all stored locally in a JSON file.

## 📦 Features

- Add a new task
- List all tasks
- Mark tasks as done
- Delete tasks
- Persistent local storage (`tasks.json`)

## 🚀 How to Use

### 1. Build the CLI
```bash
go build

2. Run commands
➕ Add a Task

./myapp add "Buy groceries"

📋 List Tasks

./myapp list

✅ Mark Task as Done

./myapp done 1

❌ Delete a Task

./myapp delete 1
