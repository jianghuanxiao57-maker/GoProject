package main

import "testing"

// 这是一个简单的测试函数
func TestMySQLConnection(t *testing.T) {
	db := initDB() // 调用你 db.go 里的初始化函数
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Errorf("数据库连接不上啊！错误: %v", err)
	}
}
