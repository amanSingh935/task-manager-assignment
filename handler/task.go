package handler

import (
	"encoding/json"
	"net/http"
	"task-manager/db"
	"task-manager/model"
	"task-manager/utils"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.Id = uuid.New().String()
	task.CreatedAt = time.Now()
	db.CreateTask(task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task model.Task

	task, exists := db.GetTask(id)
	if !exists {
		http.Error(w, "Task not present", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task model.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.Id = id
	if !db.UpdateTask(id, task) {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if !db.DeleteTask(id) {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	listResponse := utils.PrepareListResponse(r)
	page := listResponse.Page
	limit := listResponse.Limit

	allTasks := db.GetAllTasks()
	allTasks = filterTasks(allTasks, r)
	start := (page - 1) * limit // start at zero based index
	end := start + limit
	if start > len(allTasks) {
		start = len(allTasks)
	}
	if end > len(allTasks) {
		end = len(allTasks)
	}

	paginatedTasks := allTasks[start:end]

	listResponse.Total = len(allTasks)
	listResponse.PaginatedTasks = paginatedTasks

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listResponse)
}

func filterTasks(tasks []model.Task, r *http.Request) []model.Task {
	statusFilter := r.URL.Query().Get("status")
	if statusFilter == "" {
		return tasks // no filter applied
	}

	filtered := make([]model.Task, 0)
	for _, task := range tasks {
		if task.Status == statusFilter {
			filtered = append(filtered, task)
		}
	}
	return filtered
}
