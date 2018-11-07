package main

import (
	"github.com/jihuichoi/boltdb-example/db"
	"github.com/jihuichoi/boltdb-example/handler"
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
