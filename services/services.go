package services

import (
	"time"

	"github.com/optizephyr/todo-cli/models"
	"github.com/optizephyr/todo-cli/storage"
)

type Storage interface {
	LoadTasks() ([]models.Task, error)
	SaveTasks([]models.Task) error
}

var tasks []models.Task
var store Storage = storage.FileStorage{}
var loaded bool = false

func SetStorage(s Storage) {
	store = s
	loaded = false
}
func loadTasksOnce() error {
	if !loaded {
		var err error
		tasks, err = store.LoadTasks()
		if err != nil {
			return err
		}
		loaded = true
	}
	return nil
}
func saveTasks() error {
	return store.SaveTasks(tasks)
}
func AddTask(description string) (*models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return nil, err
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
	if err := saveTasks(); err != nil {
		return &task, err
	}
	return &task, nil
}

func GetAllTasks() ([]models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return []models.Task{}, err
	}
	return tasks, nil
}
func DoneTasks(ids []int) ([]models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return nil, err
	}
	updated := []models.Task{}
	for _, id := range ids {
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Done = true
				updated = append(updated, task)
			}
		}
	}
	if err := saveTasks(); err != nil {
		return updated, err
	}
	return updated, nil
}

func RemoveTasks(ids []int) ([]models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return nil, err
	}
	removed := []models.Task{}
	for _, id := range ids {
		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				removed = append(removed, task)
				break
			}
		}
	}
	if err := saveTasks(); err != nil {
		return removed, err
	}
	return removed, nil
}

func CleanTasks() ([]models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return []models.Task{}, err
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
	if err := saveTasks(); err != nil {
		return del, err
	}
	return del, nil

}

func EditTasks(id int, modified string) (*models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return nil, err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = modified
			if err := saveTasks(); err != nil {
				return &tasks[i], err
			}
		}
	}
	return nil, nil
}

func GetDoneTasks() ([]models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return []models.Task{}, err
	}
	res := []models.Task{}
	for _, task := range tasks {
		if task.Done {
			res = append(res, task)
		}
	}
	return res, nil
}
func GetPendingTasks() ([]models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return []models.Task{}, err
	}
	res := []models.Task{}
	for _, task := range tasks {
		if !task.Done {
			res = append(res, task)
		}
	}
	return res, nil
}
func UndoTasks(ids []int) ([]models.Task, error) {
	if err := loadTasksOnce(); err != nil {
		return []models.Task{}, err
	}
	res := []models.Task{}
	for _, id := range ids {
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Done = false
				res = append(res, task)
			}
		}
	}
	if err := saveTasks(); err != nil {
		return res, err
	}
	return res, nil
}
