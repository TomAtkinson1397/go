package main

import (
	"fmt"
	"log"
	"github.com/boltdb/bolt"
)
var val []byte

func getUrl(keyword string) (string, error) {
	db, err := bolt.Open("/home/tom/go/src/github.com/tomatkinson1397/goresolv/keywords.db", 0644, nil)
	logErr(err)
	defer db.Close()

	key := []byte(keyword)

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(keywords)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", keywords)
		}

		val = bucket.Get(key)
		log.Println(string(key), ": ", string(val))
		return nil
	})
	return string(val), nil
}

// func setUrl(keyword string, url string)

func logErr(err error) {
	if err != nil {
		panic(err)
	}
}
