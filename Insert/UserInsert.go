package Insert

import (
	"fmt"

	"database/sql" // 必须引入这个才能用 *sql.DB
)

// 增加 db 参数
func InsertUser(db *sql.DB, name string, age int) {
	sqlStr := "INSERT INTO users (name, age) VALUES (?, ?)"
	res, err := db.Exec(sqlStr, name, age) // 现在这里不报错了
	if err != nil {
		fmt.Println("插入失败:", err)
		return
	}
	id, _ := res.LastInsertId()
	fmt.Printf("✅ 插入成功，ID: %d\n", id)
}
