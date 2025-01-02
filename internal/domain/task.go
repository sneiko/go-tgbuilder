package domain

import "github.com/google/uuid"

// Task task for users
type Task struct {
	ID    uuid.UUID
	Title string
	Link  string
	Score float32
}

func NewTask(title, link string, score float32) *Task {
	return &Task{
		ID:    uuid.New(),
		Title: title,
		Link:  link,
		Score: score,
	}
}
