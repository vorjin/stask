package db

import (
	"time"

	"github.com/boltdb/bolt"
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
	ListToDoTasks() ([]Task, error)
	ListCompletedTasks(hours int) ([]Task, error)
	AddTask(taskName string) error
	DoTask(args []string) ([]Task, error)
	Close() error
}

type BoltTaskStore struct {
	db          *bolt.DB
	tasksBucket []byte
}

func NewBoltTaskStore(path string, tasksBucketName string) (*BoltTaskStore, error) {
	tasksBucket := []byte(tasksBucketName)

	database, err := BoltDBInit(path, tasksBucket)
	if err != nil {
		return nil, err
	}

	return &BoltTaskStore{db: database, tasksBucket: tasksBucket}, nil
}

func (s *BoltTaskStore) Close() error {
	return s.db.Close()
}
