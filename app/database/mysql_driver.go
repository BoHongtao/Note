package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

//初始化
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/base?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
