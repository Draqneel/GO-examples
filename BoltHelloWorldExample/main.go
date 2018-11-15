package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var world = []byte("world")

func main() {
	// create or open DB file if it do not exist
	db, err := bolt.Open("world.db", 0644, nil)
	// "check func on presence of errors" later "err"
	if err != nil {
		log.Fatal(err)
	}
	// when func "main" ends DB will be closed
	defer db.Close()

	// create data for record
	key := []byte("hello")
	value := []byte("Hello World!")

	// open transaction for update (read or write)
	err = db.Update(func(tx *bolt.Tx) error {
		// create bucket for recording in future (if it is not exist) + err
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		// put record in bucket
		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		// open pre-created bucket + err
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", world)
		}

		// get value from bucket by key
		val := bucket.Get(key)
		// casting from []byte to string
		fmt.Println(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
