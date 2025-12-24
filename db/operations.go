// Package db contains with operations with db and other numeric functions
package db

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var tasksBucket = []byte("tasks")
var completedBucket = []byte("completed")
var completedTimeBucket = []byte("completed_time")

func BoltDBInit(path string) error {
	var err error

	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
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

func ListToDoTasks() error {
	bucketBytes := []byte("tasks")

	return db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketBytes)

		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			if value != nil {
				id := bToU(key)
				fmt.Printf("%d. %s\n", id, value)
			}
		}

		return nil
	})
}

func ListCompletedTasks(hours int) error {
	// calculating cutoff point
	duration := time.Duration(hours) * time.Hour
	cutoff := time.Now().Add(-duration).Format(time.RFC3339)
	cutoffBytes := []byte(cutoff)

	return db.View(func(tx *bolt.Tx) error {
		timeBucket := tx.Bucket([]byte("completed_time"))
		dataBucket := tx.Bucket([]byte("completed"))

		timeCursor := timeBucket.Cursor()

		i := 1
		for key, timeValue := timeCursor.Last(); key != nil; key, timeValue = timeCursor.Prev() {
			if bytes.Compare(timeValue, cutoffBytes) < 0 {
				break
			}

			taskBytes := dataBucket.Get(key)
			if taskBytes != nil {
				// id := bToU(key)
				fmt.Printf("%d. %s\n", i, taskBytes)
				i++
			}
		}

		return nil
	})
}

func AddTask(task []byte, bucketName string) error {
	bucketBytes := []byte(bucketName)

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketBytes)

		id, err := bucket.NextSequence()
		if err != nil {
			panic(err)
		}

		idBytes := uToB(id)

		err = bucket.Put(idBytes, task)
		if err != nil {
			panic(err)
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
		panic(err)
	}

	return taskDesc, nil
}

func DeleteTask(id uint64) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(tasksBucket)

		idBytes := uToB(id)

		err := bucket.Delete(idBytes)
		if err != nil {
			panic(err)
		}

		return nil
	})
}

func DeleteTasksBucket() error {
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(tasksBucket)
	})
}

func bToU(key []byte) uint64 {
	return binary.BigEndian.Uint64(key)
}

func uToB(id uint64) []byte {
	idBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(idBytes, id)
	return idBytes
}
