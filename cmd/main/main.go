package main

import (
	"crud-server/storage"
	"crud-server/web"
	"log"
)

func main() {
	//db := storage.NewData()
	//customDataHandler := customStruct.NewDataHandler(db)
	//webDataHandler := web.NewDataHandler(customDataHandler)
	data := storage.NewDB()
	handler := web.NewDataHandler(data)
	router := web.NewPeopleStoreRouter(handler)
	port := ":8080"
	server := web.NewServer(port, router)
	println("server started, port:", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
