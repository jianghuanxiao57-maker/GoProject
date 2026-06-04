// 路径：GoProject/db/db.go
package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func InitDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Conn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
