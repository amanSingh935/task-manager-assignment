package model

import "time"

type Task struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`       // active,
	TriggerTime time.Time `json:"trigger_time"` // Regex for execution
}
