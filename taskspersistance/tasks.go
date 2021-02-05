package taskspersistance

import (
	"time"

	"github.com/boltdb/bolt"
)

var tasksBucket = []byte("tasks")
var db *bolt.DB

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(t *bolt.Tx) error {
		_, error := t.CreateBucketIfNotExists(tasksBucket)
		return error
	})
}
