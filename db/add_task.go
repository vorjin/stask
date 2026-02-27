package db

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

func (s *BoltTaskStore) AddTask(taskName string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.tasksBucket)

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		idBytes := uToB(id)

		task := Task{
			ID:           id,
			Name:         taskName,
			CreationTime: time.Now(),
		}

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
}
