package databese_sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB(password string, ip_addr string, database string) {
	db, err := sql.Open("mysql", "root:B1xSDjDAQIJjeCDJMNiM@tcp(192.168.83.6:3306)/test_go")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("database connected!")

}
