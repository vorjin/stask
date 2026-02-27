// Package db contains with operations with db and other numeric functions
package db

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

func (s *BoltTaskStore) ListToDoTasks() ([]Task, error) {
	var tasks []Task

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.tasksBucket)

		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			if value != nil {
				id := bToU(key)

				var task Task
				task.ID = id
				task.Task = string(value)

				tasks = append(tasks, task)
			}
		}

		return nil
	})

	return tasks, err
}

func (s *BoltTaskStore) ListCompletedTasks(hours int) ([]Task, error) {
	var tasks []Task

	// calculating cutoff point
	duration := time.Duration(hours) * time.Hour
	cutoff := time.Now().Add(-duration).Format(time.RFC3339)
	cutoffBytes := []byte(cutoff)

	err := s.db.View(func(tx *bolt.Tx) error {
		timeBucket := tx.Bucket(s.completedTimeBucket)
		dataBucket := tx.Bucket(s.completedBucket)

		timeCursor := timeBucket.Cursor()

		for key, timeValue := timeCursor.Last(); key != nil; key, timeValue = timeCursor.Prev() {
			if bytes.Compare(timeValue, cutoffBytes) < 0 {
				break
			}

			taskBytes := dataBucket.Get(key)
			if taskBytes != nil {
				var task Task

				id := bToU(key)
				task.ID = id
				task.Task = string(taskBytes)

				tasks = append(tasks, task)
			}
		}

		return nil
	})

	return tasks, err
}

func (s *BoltTaskStore) AddToDoTask(task []byte) error {
	return s.AddTask(task, s.tasksBucket)
}

func (s *BoltTaskStore) AddTask(task []byte, bucketBytes []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketBytes)

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		idBytes := uToB(id)

		err = bucket.Put(idBytes, task)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *BoltTaskStore) TaskByID(id uint64) ([]byte, error) {
	var taskDesc []byte

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.tasksBucket)

		idBytes := uToB(id)

		taskDesc = bucket.Get(idBytes)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return taskDesc, nil
}

func (s *BoltTaskStore) DeleteTask(id uint64) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.tasksBucket)

		idBytes := uToB(id)

		err := bucket.Delete(idBytes)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *BoltTaskStore) DeleteTasksBucket() error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(s.tasksBucket)
	})
}

func (s *BoltTaskStore) DoTask(args []string) ([]string, error) {
	var tasks []string

	for _, taskID := range args {
		id, err := strconv.ParseUint(taskID, 10, 64)
		if err != nil {
			return nil, err
		}

		taskDesc, err := s.TaskByID(id)
		if err != nil {
			return nil, err
		}

		err = s.AddTask(taskDesc, s.completedBucket)
		if err != nil {
			return nil, err
		}

		timeNow := []byte(time.Now().Format(time.RFC3339))

		err = s.AddTask(timeNow, s.completedTimeBucket)
		if err != nil {
			return nil, err
		}

		err = s.DeleteTask(id)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, taskID)
	}

	return tasks, nil
}

func bToU(key []byte) uint64 {
	return binary.BigEndian.Uint64(key)
}

func uToB(id uint64) []byte {
	idBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(idBytes, id)
	return idBytes
}
