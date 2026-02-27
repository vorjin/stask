package db

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

func (s *BoltTaskStore) DoTask(args []string) ([]Task, error) {
	var tasks []Task

	for _, taskID := range args {
		id, err := strconv.ParseUint(taskID, 10, 64)
		if err != nil {
			return nil, err
		}

		task, err := s.GetTaskByID(id)
		if err != nil {
			return nil, err
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
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *BoltTaskStore) GetTaskByID(id uint64) (Task, error) {
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
