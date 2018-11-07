package model

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

// Config type
type Config struct {
	Height   float64   `json:"height"`
	Birthday time.Time `json:"birthday"`
}

// Entry type
type Entry struct {
	Calories int    `json:"calories"`
	Food     string `json:"food"`
}

func SetConfig(db *bolt.DB, config Config) error {
	confBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("could not marshal config json: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("DB")).Put([]byte("CONFIG"), confBytes)
		if err != nil {
			return fmt.Errorf("could not set config: %v", err)
		}
		return nil
	})
	fmt.Println("Set Config")
	return err
}

func AddWeight(db *bolt.DB, weight string, date time.Time) error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("WEIGHT")).Put([]byte(date.Format(time.RFC3339)), []byte(weight))
		if err != nil {
			return fmt.Errorf("could not insert weight: %v", err)
		}
		return nil
	})
	fmt.Println("Added Weight")
	return err
}

func AddEntry(db *bolt.DB, calories int, food string, date time.Time) error {
	entry := Entry{Calories: calories, Food: food}
	entryBytes, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("could not marshal entry json: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("ENTRIES")).Put([]byte(date.Format(time.RFC3339)), entryBytes)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}

		return nil
	})
	fmt.Println("Added Entry")
	return err
}
