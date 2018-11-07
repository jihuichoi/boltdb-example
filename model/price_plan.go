package model

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

type Price struct {
	PriceByHour PriceByHour
	BasicPrice  int64
}

type PriceByHour struct {
	Time      int64
	LoadType  string
	UnitPrice int64
}

func AddPrice(db *bolt.DB, bucket string, key string, price Price) error {
	_price, _ := json.Marshal(price)
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(bucket)).Put([]byte(key), _price)
		if err != nil {
			return fmt.Errorf("could not insert weight: %v", err)
		}
		return nil
	})
	fmt.Println("Added price")
	return err
}
