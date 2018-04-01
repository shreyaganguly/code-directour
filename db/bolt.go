package db

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/boltdb/bolt"
	"github.com/shreyaganguly/code-directour/models"
)

var (
	db *bolt.DB
)

func Init(path string) error {
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

func UserExists(u string) bool {
	user, err := LookupinUser(u)
	if err != nil || user == nil {
		log.Println(err)
		return false
	}
	return true
}

//update updates key value pair in a bucket
func Update(m models.Model) error {
	bucket := m.BucketName()
	key := m.ID()
	value := m.Value()
	if err := db.Update(func(tx *bolt.Tx) error {
		if bucket == "manager" {
			b := tx.Bucket([]byte(m.BucketName()))
			var snippetInfos []*models.SnippetInfo
			var flag bool
			if err := b.ForEach(func(k, v []byte) error {
				if string(k) == key {
					err := json.Unmarshal(v, &snippetInfos)
					if err != nil {
						return err
					}
					snippetInfos = append(snippetInfos, value.(*models.SnippetInfo))
					flag = true
					return nil
				}
				return nil
			}); err != nil {
				return err
			}
			if !flag {
				snippetInfos = append(snippetInfos, value.(*models.SnippetInfo))
			}
			marshaledSnippets, err := json.Marshal(snippetInfos)
			if err != nil {
				return err
			}
			if err := tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(marshaledSnippets)); err != nil {
				return err
			}
			return nil
		}
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

func lookupinBucket(name string) (*models.SnippetInfo, error) {
	var s *models.SnippetInfo
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

func LookupinUser(name string) (*models.User, error) {
	var u *models.User
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

func All(name string) (models.Snippets, error) {
	var snippetInfos models.Snippets
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("manager"))

		// Iterate over items in sorted key order.
		if err := b.ForEach(func(k, v []byte) error {
			if string(k) == name {
				err := json.Unmarshal(v, &snippetInfos)
				if err != nil {
					return err
				}
			}
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

func Find(name, key string) (*models.SnippetInfo, error) {
	var snippetInfos []*models.SnippetInfo
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("manager"))
		// Iterate over items in sorted key order.
		if err := b.ForEach(func(k, v []byte) error {
			if string(k) == name {
				err := json.Unmarshal(v, &snippetInfos)
				if err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	for _, snippet := range snippetInfos {
		if snippet.Key == key {
			return snippet, nil
		}
	}
	return nil, errors.New(" Snippet Not found")
}

func FindAndUpdate(name, key, sharedTo string) (*models.SnippetInfo, error) {
	var snippetInfos []*models.SnippetInfo
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("manager"))
		// Iterate over items in sorted key order.
		if err := b.ForEach(func(k, v []byte) error {
			if string(k) == name {
				err := json.Unmarshal(v, &snippetInfos)
				if err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	var sharedSnippet *models.SnippetInfo
	for _, snippet := range snippetInfos {
		if snippet.Key == key {
			sharedSnippet = snippet
			snippet.SharedToSomeone = true
			snippet.SharedTo = sharedTo
		}

	}
	marshaledSnippets, err := json.Marshal(snippetInfos)
	if err != nil {
		return nil, err
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket([]byte("manager")).Put([]byte(name), []byte(marshaledSnippets)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return sharedSnippet, err
	}
	return sharedSnippet, nil
}

func Delete(name, key string) error {
	var snippetInfos []*models.SnippetInfo
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("manager"))
		// Iterate over items in sorted key order.
		if err := b.ForEach(func(k, v []byte) error {
			if string(k) == name {
				err := json.Unmarshal(v, &snippetInfos)
				if err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	var flag bool
	for i, snippet := range snippetInfos {
		if snippet.Key == key && !flag {
			snippetInfos = append(snippetInfos[:i], snippetInfos[i+1:]...)
			flag = true
		}
	}
	marshaledSnippets, err := json.Marshal(snippetInfos)
	if err != nil {
		return err
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket([]byte("manager")).Put([]byte(name), []byte(marshaledSnippets)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
