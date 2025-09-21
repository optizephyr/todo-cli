package services

import (
	"fmt"
	"time"

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
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	storage.SaveTasks(tasks)
	fmt.Printf("Added task %d: %s\n", task.ID, task.Description)
}

func ListAllTasks() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("%d. [%v] %s created at %v\n", task.ID, task.Done, task.Description, task.CreatedAt)
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
			fmt.Printf("task %d is removed\n", id)
			return
		}
	}
	fmt.Printf("task %d not found\n", id)
}

func CleanTasks() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}
	res := []models.Task{}
	del := []models.Task{}
	for _, task := range tasks {
		if task.Done {
			del = append(del, task)
		} else {
			res = append(res, task)
		}
	}
	if len(del) == 0 {
		fmt.Println("Nothing to do")
		return
	}
	fmt.Println("Clean the below tasks:")
	for _, task := range del {
		fmt.Printf("%d. %s created at %v\n", task.ID, task.Description, task.CreatedAt)
	}
	storage.SaveTasks(res)
}

func EditTasks(id int, modified string) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = modified
			fmt.Printf("task %d become %s\n", id, modified)
			storage.SaveTasks(tasks)
			return
		}
	}
	fmt.Printf("task %d not found\n", id)
}

func ListTasksWithExp(done bool) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}

	flag := true
	for _, task := range tasks {
		if task.Done == done {
			fmt.Printf("%d. [%v] %s created at %v\n", task.ID, task.Done, task.Description, task.CreatedAt)
			flag = false
		}
	}
	if flag {
		fmt.Println("No tasks found.")
	}
}

func ListDoneTasks() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}

	flag := true
	for _, task := range tasks {
		if task.Done {
			fmt.Printf("%d. [%v] %s created at %v\n", task.ID, task.Done, task.Description, task.CreatedAt)
			flag = false
		}
	}
	if flag {
		fmt.Println("No tasks found.")
	}
}
func ListPendingTasks() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}

	flag := true
	for _, task := range tasks {
		if !task.Done {
			fmt.Printf("%d. [%v] %s created at %v\n", task.ID, task.Done, task.Description, task.CreatedAt)
			flag = false
		}
	}
	if flag {
		fmt.Println("No tasks found.")
	}
}
