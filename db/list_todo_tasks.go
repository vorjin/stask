package db

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

func (s *BoltTaskStore) ListToDoTasks() ([]Task, error) {
	var tasks []Task

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.tasksBucket)

		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			var task Task
			err := json.Unmarshal(value, &task)
			if err != nil {
				return err
			}

			if task.CompletionTime.IsZero() && task.DeletionTime.IsZero() {
				tasks = append(tasks, task)
			}
		}

		return nil
	})

	return tasks, err
}
