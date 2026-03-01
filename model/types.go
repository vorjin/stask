// Package model defines the data structures and interfaces for the task management system.
package model

import "time"

type TaskStatus int

const (
	Todo TaskStatus = iota
	Completed
	Deleted
)

type Task struct {
	ID             uint64
	Name           string
	Description    string
	CreationTime   time.Time
	CompletionTime time.Time
	DeletionTime   time.Time
}

type TaskStore interface {
	ListTasks(status TaskStatus, hours int) ([]Task, error)
	AddTask(taskName string) error
	UpdateTask(status TaskStatus, taskID uint64) (Task, error)
	Close() error
}
