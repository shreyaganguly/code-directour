package main

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
)

var (
	db *bolt.DB
)

func initDB(path string) error {
	var err error
	db, err = bolt.Open(path, 0644, nil)
	if err != nil {
		return err
	}
	return createBucket("manager", "user")
}

func createBucket(names ...string) error {
	for _, name := range names {
		if err := db.Update(func(tx *bolt.Tx) error {
			var err error
			_, err = tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

//update updates key value pair in a bucket
func update(key, bucket string, value interface{}) error {
	if err := db.Update(func(tx *bolt.Tx) error {
		marshaledValue, err := json.Marshal(value)

		if err != nil {
			return err
		}
		if err := tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(marshaledValue)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func lookupinBucket(name string) (*SnippetInfo, error) {
	var s *SnippetInfo
	e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("manager"))

		if b == nil {
			return errors.New("bucket not found")
		}

		v := b.Get([]byte(name))

		err := json.Unmarshal(v, &s)

		if err != nil {
			return err
		}

		return nil
	})

	return s, e
}

func lookupinUser(name string) (*User, error) {
	var u *User
	e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))

		if b == nil {
			return errors.New("bucket not found")
		}

		v := b.Get([]byte(name))

		err := json.Unmarshal(v, &u)

		if err != nil {
			return err
		}

		return nil
	})

	return u, e
}

func all() ([]*SnippetInfo, error) {
	var snippetInfos []*SnippetInfo
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("manager"))

		// Iterate over items in sorted key order.
		if err := b.ForEach(func(k, v []byte) error {
			var s *SnippetInfo
			err := json.Unmarshal(v, &s)
			if err != nil {
				return err
			}
			snippetInfos = append(snippetInfos, s)
			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return snippetInfos, nil
}
