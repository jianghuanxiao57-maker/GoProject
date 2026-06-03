package Query

import (
	"fmt"

	"database/sql" // 必须引入这个才能用 *sql.DB
)

func QueryUser(db *sql.DB, id int) {
	var name string
	var age int
	err := db.QueryRow("SELECT name, age FROM users WHERE id = ?", id).Scan(&name, &age)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	fmt.Printf("🔍 查到用户: %s, 年龄: %d\n", name, age)
}
