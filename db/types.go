package db

import "github.com/boltdb/bolt"

type Task struct {
	ID   uint64
	Task string
}

type TaskStore interface {
	ListToDoTasks() ([]Task, error)
	ListCompletedTasks(hours int) ([]Task, error)
	AddToDoTask(task []byte) error
	AddTask(task []byte, bucketName []byte) error
	TaskByID(id uint64) ([]byte, error)
	DeleteTask(id uint64) error
	DeleteTasksBucket() error
	DoTask(args []string) ([]string, error)
	Close() error
}

type BoltTaskStore struct {
	db                  *bolt.DB
	tasksBucket         []byte
	completedBucket     []byte
	completedTimeBucket []byte
}

func NewBoltTaskStore(path string, tasksBucketName string, completedBucketName string, completedTimeBucketName string) (*BoltTaskStore, error) {
	tasksBucket := []byte(tasksBucketName)
	completedBucket := []byte(completedBucketName)
	completedTimeBucket := []byte(completedTimeBucketName)

	database, err := BoltDBInit(path, tasksBucket, completedBucket, completedTimeBucket)
	if err != nil {
		return nil, err
	}

	return &BoltTaskStore{db: database, tasksBucket: tasksBucket, completedBucket: completedBucket, completedTimeBucket: completedTimeBucket}, nil
}

func (s *BoltTaskStore) Close() error {
	return s.db.Close()
}
