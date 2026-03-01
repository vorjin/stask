package db

import (
	"encoding/json"
	"errors"
	"time"

	"stask/model"

	"github.com/boltdb/bolt"
)

func (s *BoltTaskStore) ListTasks(status model.TaskStatus, hours int) ([]model.Task, error) {
	var tasks []model.Task

	// calculating cutoff point
	duration := time.Duration(hours) * time.Hour
	cutoff := time.Now().Add(-duration)

	// db operataion
	err := s.db.View(func(tx *bolt.Tx) error {
		tasksBucket := tx.Bucket(s.tasksBucket)
		cursor := tasksBucket.Cursor()

		for key, value := cursor.Last(); key != nil; key, value = cursor.Prev() {
			var task model.Task
			err := json.Unmarshal(value, &task)
			if err != nil {
				return err
			}

			switch status {
			case model.Todo:
				if task.CompletionTime.IsZero() && task.DeletionTime.IsZero() {
					tasks = append(tasks, task)
				}
			case model.Completed:
				if task.CompletionTime.After(cutoff) {
					tasks = append(tasks, task)
				}
			case model.Deleted:
				if task.DeletionTime.After(cutoff) {
					tasks = append(tasks, task)
				}
			default:
				return errors.New("invalid task status")
			}
		}

		return nil
	})

	return tasks, err
}
