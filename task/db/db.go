package db

import (
	"encoding/binary"
	"encoding/json"
	"log"

	"github.com/bolt"
)

var taskBucket []byte = []byte("Tasks")
var db *bolt.DB

type Task struct {
	Id   uint64
	Text string
	Done bool
}

func init() {
	var err error
	db, err = bolt.Open("task.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
}

func DoTask(id uint64) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		bytes := b.Get(itob(id))
		var task Task
		err := json.Unmarshal(bytes, &task)
		if err != nil {
			return err
		}
		task.Done = true
		tBytes, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return b.Put(itob(id), tBytes)
	})
}

func AddTask(t string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id, _ := b.NextSequence()
		task, err := json.Marshal(Task{
			Id:   id,
			Text: t,
			Done: false,
		})
		if err != nil {
			return err
		}
		return b.Put(itob(id), task)
	})
}

func ListTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		b.ForEach(func(k, v []byte) error {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}
			if !task.Done {
				tasks = append(tasks, task)
			}
			return nil
		})
		return nil
	})
	return tasks, err
}

func itob(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	return b
}
