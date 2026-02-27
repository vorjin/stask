package db

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
)

func (s *BoltTaskStore) DoTask(taskID uint64) (Task, error) {
	task, err := s.getTaskByID(taskID)
	if err != nil {
		return Task{}, err
	}

	task.CompletionTime = time.Now()

	err = s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.tasksBucket)

		idBytes := uToB(task.ID)
		taskBytes, err := json.Marshal(task)
		if err != nil {
			return err
		}

		err = bucket.Put(idBytes, taskBytes)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *BoltTaskStore) getTaskByID(id uint64) (Task, error) {
	var task Task

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.tasksBucket)

		idBytes := uToB(id)

		taskBytes := bucket.Get(idBytes)
		if taskBytes == nil {
			return errors.New("task not found")
		}

		err := json.Unmarshal(taskBytes, &task)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return Task{}, err
	}

	return task, nil
}
