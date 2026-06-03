package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // 必须引入驱动，注意前面的下划线
)

func main() {
	// 连接字符串格式：用户名:密码@tcp(地址:端口)/数据库名
	dsn := "root:123456@tcp(127.0.0.1:3306)/my_db"

	// 1. 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer db.Close() // 确保程序退出时关闭连接

	// 2. 尝试测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库无法访问:", err)
	}

	fmt.Println("数据库连接成功！")

	// 3. 设置连接池（这一步是 Go 比 Java 方便的地方，直接内置了）
	db.SetMaxOpenConns(10)                 // 最大连接数
	db.SetConnMaxLifetime(time.Minute * 3) // 连接最大生命周期
}
