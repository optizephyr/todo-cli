package services

import (
	"fmt"
	"time"

	"github.com/optizephyr/todo-cli/models"
	"github.com/optizephyr/todo-cli/storage"
)

var tasks []models.Task
var loaded bool = false

func loadTasksOnce() error {
	if !loaded {
		var err error
		tasks, err = storage.LoadTasks()
		if err != nil {
			return err
		}
		loaded = true
	}
	return nil
}
func saveTasks() error {
	return storage.SaveTasks(tasks)
}
func AddTask(description string) {
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
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
	fmt.Printf("Added task %d: %s\n", task.ID, task.Description)
	if err := saveTasks(); err != nil {
		fmt.Println(err)
	}
}

func ListAllTasks() {
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("%d. [%v] %s created at %v\n", task.ID, task.Done, task.Description, task.CreatedAt)
	}
}
func DoneTasks(ids []int) {
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
	}
	for _, id := range ids {
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Done = true
				fmt.Printf("task %d have done\n", id)
				if err := saveTasks(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func RemoveTasks(ids []int) {
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
	}
	for _, id := range ids {
		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				fmt.Printf("task %d is removed\n", id)
				if err := saveTasks(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func CleanTasks() {
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
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
	tasks = res
	if len(del) == 0 {
		fmt.Println("Nothing to do")
		return
	}
	fmt.Println("Clean the below tasks:")
	for _, task := range del {
		fmt.Printf("%d. %s created at %v\n", task.ID, task.Description, task.CreatedAt)
	}
	if err := saveTasks(); err != nil {
		fmt.Println(err)
	}
}

func EditTasks(id int, modified string) {
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = modified
			fmt.Printf("task %d become %s\n", id, modified)
			if err := saveTasks(); err != nil {
				fmt.Println(err)
			}
			return
		}
	}
	fmt.Printf("task %d not found\n", id)
}

func ListDoneTasks() {
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
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
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
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
func UndoTasks(ids []int) {
	if err := loadTasksOnce(); err != nil {
		fmt.Println(err)
		return
	}
	for _, id := range ids {
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Done = false
				fmt.Printf("task %d have done\n", id)
				if err := saveTasks(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
