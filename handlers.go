package main

import (
	"fmt"
	"github.com/optizephyr/todo-cli/models"
	"github.com/optizephyr/todo-cli/storage"
)

func AddTask(description string) {
	maxID := 0
	for _, task := range tasks {
		if maxID < task.ID {
			maxID = task.ID
		}
	}
	task := models.Task{
		ID:          maxID + 1,
		Description: description,
		Done:        false,
	}

	tasks = append(tasks, task)
	storage.SaveTasks(tasks)
	fmt.Printf("Added task %d: %s\n", task.ID, task.Description)
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("%d. [%v] %s\n", task.ID, task.Done, task.Description)
	}
}
func DoneTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			return
		}
	}
	storage.SaveTasks(tasks)
	fmt.Printf("task %d not found\n", id)
}

func RemoveTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
	storage.SaveTasks(tasks)
	fmt.Printf("task %d not found\n", id)
}
