package db

import (
	"github.com/boltdb/bolt"
)

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
