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
			id := uToI(key)
			fmt.Printf("%d. %s\n", id, value)
		}

		return nil
	})
}

func uToI(key []byte) uint64 {
	return binary.BigEndian.Uint64(key)
}
