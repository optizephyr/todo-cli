package services

import (
	"fmt"

	"github.com/optizephyr/todo-cli/models"
	"github.com/optizephyr/todo-cli/storage"
)

func AddTask(description string) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}
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
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("%d. [%v] %s\n", task.ID, task.Done, task.Description)
	}
}
func DoneTask(id int) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Printf("task %d have done\n", id)
			storage.SaveTasks(tasks)
			return
		}
	}
	fmt.Printf("task %d not found\n", id)
}

func RemoveTask(id int) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			storage.SaveTasks(tasks)
			return
		}
	}
	fmt.Printf("task %d not found\n", id)
}
