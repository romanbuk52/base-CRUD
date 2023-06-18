package database_sql

import (
	"crud-server/web"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	sql *sql.DB
}

// DROP TABLE IF EXISTS album;
// CREATE TABLE album (
//   id         INT AUTO_INCREMENT NOT NULL,
//   title      VARCHAR(128) NOT NULL,
//   artist     VARCHAR(255) NOT NULL,
//   price      DECIMAL(5,2) NOT NULL,
//   PRIMARY KEY (`id`)
// );

func ConnectDB(ip_addr string, database string, user string, password string) *Database {
	db, err := sql.Open("mysql", (user + ":" + password + "@tcp" + "(\\ip_addr)" + "/" + database))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	fmt.Println("database connected!")
	data := &Database{
		sql: db,
	}
	return data
}

// get MultiQuery
func (dtb *Database) GetAll() {
}

func (dtb *Database) Get(id string) (m web.Man, er error) {
	dtb.sql.Query()
}

func (dtb *Database) Add() {
	dtb.sql.Exec()
}

func (dtb *Database) Edit(m web.Man) {

}

func (dtb *Database) Del(id string) error {

}
