package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server_port string
	DB_address  string
	DB_name     string
	DB_user     string
	DB_pass     string
}

func getConf(conf_path string) *Config {
	yamlFile, err := ioutil.ReadFile(conf_path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var c Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &c
}
