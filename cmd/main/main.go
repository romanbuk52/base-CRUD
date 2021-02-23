package main

import (
	"crud-server/web"
	"log"
)

func main() {
	//db := storage.NewData()
	//customDataHandler := customStruct.NewDataHandler(db)
	//webDataHandler := web.NewDataHandler(customDataHandler)
	handler := web.DataHandler{}
	router := web.NewPeopleStoreRouter(&handler)
	port := ":8080"
	server := web.NewServer(port, router)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	println("server started, port:", port)
}
