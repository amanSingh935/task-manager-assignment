package db

import "task-manager/model"

var (
	tasks = make(map[string]model.Task)
)

func CreateTask(task model.Task) {
	tasks[task.Id] = task
}

func GetTask(id string) (model.Task, bool) {
	task, exists := tasks[id]
	return task, exists
}

func UpdateTask(id string, task model.Task) bool {
	if _, exists := tasks[id]; !exists {
		return false
	}
	tasks[id] = task
	return true
}

func DeleteTask(id string) bool {
	if _, exists := tasks[id]; !exists {
		return false
	}
	delete(tasks, id)
	return true
}

func GetAllTasks() []model.Task {
	result := make([]model.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, task)
	}
	return result
}
