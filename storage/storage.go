package storage

import (
	"encoding/json"
	"fmt"
	"github.com/optizephyr/todo-cli/models"
	"os"
)

func LoadTasks() ([]models.Task, error) {
	jsonData, err := os.ReadFile(".todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(".todos.json")
			if err != nil {
				return nil, fmt.Errorf("创建.todos.json失败:%v", err)
			}
			defer file.Close()
			return []models.Task{}, nil
		}
		return nil, fmt.Errorf("读取.todos.json失败:%v", err)
	}
	if len(jsonData) == 0 {
		return []models.Task{}, nil
	}
	var tasks []models.Task
	err = json.Unmarshal([]byte(jsonData), &tasks)
	if err != nil {
		return nil, fmt.Errorf("unmarshal失败:%v", err)
	}
	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("marshal失败:%v", err)
	}

	err = os.WriteFile(".todos.json", []byte(jsonData), 0644)
	if err != nil {
		return fmt.Errorf("写入.todos.json失败:%v", err)
	}
	return nil
}
