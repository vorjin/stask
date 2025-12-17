// Package db contains with operations with db and other numeric functions
package db

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

var db *bolt.DB
var bucketPath = []byte("tasks")

func BoltDBInit(path string) error {
	var err error

	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketPath)
		return err
	})
}

func ListTasks() error {
	fmt.Printf("This is your tasks:\n")

	return db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketPath)

		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			id := bToU(key)
			fmt.Printf("%d. %s\n", id, value)
		}

		return nil
	})
}

func AddTask(task []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketPath)

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

func DoTask(id uint64) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketPath)

		idBytes := uToB(id)

		err := bucket.Delete(idBytes)
		if err != nil {
			panic(err)
		}

		return nil
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

