package db

import (
	"time"

	"github.com/boltdb/bolt"
)

func BoltDBInit(path string, tasksBucket []byte) (*bolt.DB, error) {
	db, err := bolt.Open(path, 0o600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
