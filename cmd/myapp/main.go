package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func loadTasks() ([]Task, error) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func addTask(taskName string) {
	tasks, _ := loadTasks()

	newTask := Task{Name: taskName, Done: false}
	tasks = append(tasks, newTask)

	err := saveTasks(tasks)
	if err != nil {
		fmt.Println("error saving task:", err)
		return
	}
	fmt.Println("task added:", taskName)

}

func listTasks() {
	tasks, err := loadTasks()
	if err != nil || len(tasks) == 0 {
		fmt.Println("no tasks found.")
		return
	}

	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[X]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Name)
	}
}

func markTaskAsDone(taskNumber int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if taskNumber > len(tasks) {
		fmt.Println("Task number out of range.")
		return
	}

	tasks[taskNumber-1].Done = true

	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Marked task as done:", tasks[taskNumber-1].Name)
}

func deleteTask(taskNumber int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if taskNumber > len(tasks) {
		fmt.Println("Task number out of range.")
		return
	}

	tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)

	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task deleted.")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a task")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task name.")
			return
		}
		addTask(os.Args[2])

	case "list":

		listTasks()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task number to mark as done.")
			return
		}
		taskNumStr := os.Args[2]
		taskNumber, err := strconv.Atoi(taskNumStr)
		if err != nil || taskNumber < 1 {
			fmt.Println("Invalid Task number.")
			return
		}

		markTaskAsDone(taskNumber)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("please provide a task number to delete.")
			return
		}

		taskNumStr := os.Args[2]
		taskNumber, err := strconv.Atoi(taskNumStr)
		if err != nil || taskNumber < 1 {
			fmt.Println("Please enter a valid task nuimber to delete.")
			return
		}

		deleteTask(taskNumber)

	default:
		fmt.Println("Unknown command. Available commands are: add, list, done, delete")

	}

}
