package main

import (
	"GoProject/Query"
)

func main() {
	db := initDB() // 先初始化数据库
	defer db.Close()
	// Insert.InsertUser("李四", 30) // 调用插入逻辑
	Query.QueryUser(db, 1) // 调用查询逻辑
}
