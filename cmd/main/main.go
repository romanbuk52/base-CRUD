package main

import (
	"crud-server/storageSQL"
	"crud-server/web"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("storageSQL/storage2.db"))
	if err != nil {
		panic("failed to connect database")
	}
	// data := storage.NewDB()
	data := storageSQL.NewDB(db)
	handler := web.NewDataHandler(data)
	router := web.NewPeopleStoreRouter(handler)
	port := ":8080"
	server := web.NewServer(port, router)
	println("server started, port:", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
