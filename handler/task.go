package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-manager/db"
	"task-manager/model"
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
	page := 1
	limit := 10

	if p := r.URL.Query().Get("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}

	allTasks := db.GetAllTasks()

	start := (page - 1) * limit
	end := start + limit
	if start > len(allTasks) {
		start = len(allTasks)
	}
	if end > len(allTasks) {
		end = len(allTasks)
	}

	paginatedTasks := allTasks[start:end]

	response := map[string]interface{}{
		"page":       page,
		"limit":      limit,
		"total":      len(allTasks),
		"totalTasks": (len(allTasks) + limit - 1) / limit,
		"tasks":      paginatedTasks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
