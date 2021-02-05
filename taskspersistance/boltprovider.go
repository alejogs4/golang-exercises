package taskspersistance

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

type BoltProvider struct{}

func (b *BoltProvider) CreateTask(task string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		id64, _ := b.NextSequence()
		id := int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return err
	}
	return nil
}

func (b *BoltProvider) ListAllTasks() ([]Task, error) {
	var tasks []Task

	error := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(tasksBucket)
		tasksCursor := bucket.Cursor()
		for key, task := tasksCursor.First(); key != nil; key, task = tasksCursor.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(key),
				Value: string(task),
			})
		}

		return nil
	})

	if error != nil {
		return nil, error
	}

	return tasks, nil
}

func (b *BoltProvider) DeleteTask(taskID int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(tasksBucket)
		return bucket.Delete(itob(taskID))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
