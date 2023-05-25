package main

import (
	"crud-server/storage_gorm"
	"crud-server/usersstorage"
	"crud-server/web"
	"flag"
	"io/ioutil"
	"log"
	"runtime"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	runtime.GOMAXPROCS(4)
	pathDB := flag.String("pathDB", "storageSQL/storage0.db", "Input path to database")
	pathDBU := flag.String("pathDBU", "usersstorage/storageUsers.db", "Input path to database Users")
	pathToConfig := flag.String("config", "configs/config_DB.yml", "Input path to configuration file")
	port := flag.String("port", ":8081", "port of server")
	flag.Parse()
	var c Conf
	c.getConf(*pathToConfig)

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
	data := storage_gorm.NewDB(db)
	handler := web.NewDataHandler(data)
	router := web.NewPeopleStoreRouter(handler, handlerUsers)
	server := web.NewServer(*port, router)
	println("server started, port:", *port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

type Conf struct {
	server_port string
	DB_address  string
	DB_name     string
	DB_user     string
	DB_pass     string
}

func (c *Conf) getConf(conf_path string) *Conf {
	yamlFile, err := ioutil.ReadFile("\\conf_path")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
