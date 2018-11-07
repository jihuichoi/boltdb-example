package handler

import (
	"fmt"
	"github.com/boltdb/bolt"

	"github.com/zupzup/boltdb-example/model"
)

func ABC(db *bolt.DB) {

	bucket := "플랜1"

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})

	price := model.Price{BasicPrice: 100}}
	aa := model.PriceByHour{LoadType: "경부하", UnitPrice: 50, Time: 10}
	price.PriceByHour = append(price.PriceByHour, newPrice)

	err = model.AddPrice(db, bucket, "11", price)
	if err != nil {
		log.Fatal(err)
	}
	price = model.Price{PriceByHour: model.PriceByHour{LoadType: "경부하", UnitPrice: 50, Time: 10}, BasicPrice: 100}
	err = model.AddPrice(db, bucket, "12", price)
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// conf := model.Config{Height: 186.0, Birthday: time.Now()}
	// err := model.SetConfig(db, conf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = model.AddWeight(db, "85.0", time.Now())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = model.AddEntry(db, 100, "apple", time.Now())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// err = model.AddEntry(db, 100, "orange", time.Now().AddDate(0, 0, -2))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// err = db.View(func(tx *bolt.Tx) error {
	// 	conf := tx.Bucket([]byte("DB")).Get([]byte("CONFIG"))
	// 	fmt.Printf("Config: %s\n", conf)
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// err = db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("DB")).Bucket([]byte("WEIGHT"))
	// 	b.ForEach(func(k, v []byte) error {
	// 		fmt.Println(string(k), string(v))
	// 		return nil
	// 	})
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// err = db.View(func(tx *bolt.Tx) error {
	// 	c := tx.Bucket([]byte("DB")).Bucket([]byte("ENTRIES")).Cursor()
	// 	min := []byte(time.Now().AddDate(0, 0, -7).Format(time.RFC3339))
	// 	max := []byte(time.Now().AddDate(0, 0, 0).Format(time.RFC3339))
	// 	for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
	// 		fmt.Println(string(k), string(v))
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
