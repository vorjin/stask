// Package db contains with operations with db and other numeric functions
package db

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db                  *bolt.DB
	tasksBucket         = []byte("tasks")
	completedBucket     = []byte("completed")
	completedTimeBucket = []byte("completed_time")
)

type Task struct {
	ID   uint64
	Task string
}

func BoltDBInit(path string) error {
	var err error

	db, err = bolt.Open(path, 0o600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(completedBucket)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(completedTimeBucket)
		return err
	})
}

func ListToDoTasks() ([]Task, error) {
	bucketBytes := []byte("tasks")
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketBytes)

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

func ListCompletedTasks(hours int) ([]Task, error) {
	var tasks []Task

	// calculating cutoff point
	duration := time.Duration(hours) * time.Hour
	cutoff := time.Now().Add(-duration).Format(time.RFC3339)
	cutoffBytes := []byte(cutoff)

	err := db.View(func(tx *bolt.Tx) error {
		timeBucket := tx.Bucket([]byte("completed_time"))
		dataBucket := tx.Bucket([]byte("completed"))

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

func AddToDoTask(task []byte) error {
	return AddTask(task, "tasks")
}

func AddTask(task []byte, bucketName string) error {
	bucketBytes := []byte(bucketName)

	return db.Update(func(tx *bolt.Tx) error {
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

func TaskByID(id uint64) ([]byte, error) {
	var taskDesc []byte

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(tasksBucket)

		idBytes := uToB(id)

		taskDesc = bucket.Get(idBytes)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return taskDesc, nil
}

func DeleteTask(id uint64) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(tasksBucket)

		idBytes := uToB(id)

		err := bucket.Delete(idBytes)
		if err != nil {
			return err
		}

		return nil
	})
}

func DeleteTasksBucket() error {
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(tasksBucket)
	})
}

func DoTask(args []string) ([]string, error) {
	var tasks []string

	for _, taskID := range args {
		id, err := strconv.ParseUint(taskID, 10, 64)
		if err != nil {
			return nil, err
		}

		taskDesc, err := TaskByID(id)
		if err != nil {
			return nil, err
		}

		err = AddTask(taskDesc, "completed")
		if err != nil {
			return nil, err
		}

		timeNow := []byte(time.Now().Format(time.RFC3339))

		err = AddTask(timeNow, "completed_time")
		if err != nil {
			return nil, err
		}

		err = DeleteTask(id)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, taskID)
	}

	return tasks, nil
}

func CloseBoltDB() error {
	return db.Close()
}

func bToU(key []byte) uint64 {
	return binary.BigEndian.Uint64(key)
}

func uToB(id uint64) []byte {
	idBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(idBytes, id)
	return idBytes
}
