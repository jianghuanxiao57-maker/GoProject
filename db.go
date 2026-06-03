package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局变量
var db *sql.DB

func initDB() *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 💡 必须加上这一行，把连接好的 db 对象返回给调用者
	return db
}
