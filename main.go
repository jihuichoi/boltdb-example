package main

import (
	"github.com/zupzup/boltdb-example/db"
	"github.com/zupzup/boltdb-example/handler"
	"log"
)

func main() {
	db, err := db.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler.ABC(db)

}
