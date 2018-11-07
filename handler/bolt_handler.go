package handler

import (
	"fmt"
	"go.etcd.io/bbolt"

	"github.com/jihuichoi/boltdb-example/model"
	"log"
)

func ABC(db *bolt.DB) {

	bucket := "일반/산업용 을 고압A 선택III"
	month := "11"
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})

	price := model.Price{BasicPrice: 9810}
	price.PriceByHours = map[string]model.PriceByHour{}
	price.PriceByHours["0"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["1"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["2"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["3"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["4"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["5"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["6"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["7"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["8"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}
	price.PriceByHours["9"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 108.6}
	price.PriceByHours["10"] = model.PriceByHour{LoadType: "최대부하", UnitPrice: 155.5}
	price.PriceByHours["11"] = model.PriceByHour{LoadType: "최대부하", UnitPrice: 155.5}
	price.PriceByHours["12"] = model.PriceByHour{LoadType: "중간부하", UnitPrice: 108.6}
	price.PriceByHours["13"] = model.PriceByHour{LoadType: "중간부하", UnitPrice: 108.6}
	price.PriceByHours["14"] = model.PriceByHour{LoadType: "중간부하", UnitPrice: 108.6}
	price.PriceByHours["15"] = model.PriceByHour{LoadType: "중간부하", UnitPrice: 108.6}
	price.PriceByHours["16"] = model.PriceByHour{LoadType: "중간부하", UnitPrice: 108.6}
	price.PriceByHours["17"] = model.PriceByHour{LoadType: "최대부하", UnitPrice: 155.5}
	price.PriceByHours["18"] = model.PriceByHour{LoadType: "최대부하", UnitPrice: 155.5}
	price.PriceByHours["19"] = model.PriceByHour{LoadType: "최대부하", UnitPrice: 155.5}
	price.PriceByHours["20"] = model.PriceByHour{LoadType: "중간부하", UnitPrice: 108.6}
	price.PriceByHours["21"] = model.PriceByHour{LoadType: "중간부하", UnitPrice: 108.6}
	price.PriceByHours["22"] = model.PriceByHour{LoadType: "최대부하", UnitPrice: 155.5}
	price.PriceByHours["23"] = model.PriceByHour{LoadType: "경부하", UnitPrice: 62.5}

	// PriceByHour: model.PriceByHour{LoadType: "경부하", UnitPrice: 50, Time: 10},
	err = model.AddPrice(db, bucket, month, price)
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
