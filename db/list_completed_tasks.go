package db

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

func (s *BoltTaskStore) ListCompletedTasks(hours int) ([]Task, error) {
	var tasks []Task

	// calculating cutoff point
	duration := time.Duration(hours) * time.Hour
	cutoff := time.Now().Add(-duration)

	err := s.db.View(func(tx *bolt.Tx) error {
		tasksBucket := tx.Bucket(s.tasksBucket)
		cursor := tasksBucket.Cursor()

		for key, value := cursor.Last(); key != nil; key, value = cursor.Prev() {
			var task Task
			err := json.Unmarshal(value, &task)
			if err != nil {
				return err
			}

			if task.CompletionTime.After(cutoff) {
				tasks = append(tasks, task)
			}
		}

		return nil
	})

	return tasks, err
}
