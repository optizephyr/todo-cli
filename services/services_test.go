package services

import (
	"testing"

	"github.com/optizephyr/todo-cli/models"
)

type mockStorage struct{}

func (s mockStorage) LoadTasks() ([]models.Task, error) {
	return []models.Task{}, nil
}
func (s mockStorage) SaveTasks([]models.Task) error {
	return nil
}

func TestAddTask(t *testing.T) {
	SetStorage(mockStorage{})
	desc := "test1"
	task, err := AddTask(desc)
	if err != nil {
		t.Errorf("新增任务失败:%v\n", err)
	} else if task.Description != desc || task.Done || task.ID != 1 {
		t.Errorf("期望：%s,实际：%s\n", desc, task.Description)
	}
}

func TestCleanTasks(t *testing.T) {
	SetStorage(mockStorage{})
	AddTask("test1")
	AddTask("test2")
	AddTask("test3")
	DoneTasks([]int{1, 3})
	tasks, err := CleanTasks()
	if err != nil {
		t.Errorf("清理已完成任务失败:%v\n", err)
	} else {
		if len(tasks) != 2 || tasks[0].ID != 1 || tasks[1].ID != 3 {
			t.Errorf("期望：")
		}
	}
}
