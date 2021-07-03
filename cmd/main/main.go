package main

import (
	"crud-server/storageSQL"
	"crud-server/usersstorage"
	"crud-server/web"
	"flag"
	"log"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	runtime.GOMAXPROCS(4)
	pathDB := flag.String("pathDB", "storageSQL/storage0.db", "Input path to database")
	pathDBU := flag.String("pathDBU", "usersstorage/storageUsers.db", "Input path to database Users")
	port := flag.String("port", ":8081", "port of server")
	flag.Parse()

	db, err0 := gorm.Open(sqlite.Open(*pathDB))
	if err0 != nil {
		log.Fatalf("failed to connect database: %s", err0)
	}

	// dbUsers, pathDBU - this is a database users service
	dbUsers, err1 := gorm.Open(sqlite.Open(*pathDBU))
	if err1 != nil {
		log.Fatalf("failed to connect database: %s", err1)
	}
	dataUsers := usersstorage.NewDBU(dbUsers)
	handlerUsers := usersstorage.NewDBUHandler(dataUsers)

	// data := storage.NewDB()
	data := storageSQL.NewDB(db)
	handler := web.NewDataHandler(data)
	router := web.NewPeopleStoreRouter(handler, handlerUsers)
	server := web.NewServer(*port, router)
	println("server started, port:", *port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
