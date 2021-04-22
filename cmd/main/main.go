package main

import (
	"crud-server/storageSQL"
	"crud-server/web"
	"flag"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	pathDB := flag.String("pathDB", "/home/roman/go/src/crud-server/storageSQL/storage0.db", "Input path to database")
	port := flag.String("port", ":8080", "port of server")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open(*pathDB))
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	// data := storage.NewDB()
	data := storageSQL.NewDB(db)
	handler := web.NewDataHandler(data)
	router := web.NewPeopleStoreRouter(handler)
	server := web.NewServer(*port, router)
	println("server started, port:", *port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
