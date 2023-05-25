package databese_sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	sql *sql.DB
}

func connectDB(ip_addr string, database string, user string, password string) *Database {
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

func (dtb *Database) GetRec() {

}

func (dtb *Database) WriteRec() {

}

func (dtb *Database) UpdateRec() {

}

func (dtb *Database) DelRec() {

}
