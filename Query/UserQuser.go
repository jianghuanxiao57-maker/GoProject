package Query

import (
	"database/sql"
)

// User 定义返回的数据结构
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 修改函数，使其返回数据和错误，而不是直接 fmt.Println
func QueryUserList(db *sql.DB, nameFilter string, page int, pageSize int) ([]User, error) {
	// 1. 计算分页偏移量
	// SQL 公式: LIMIT pageSize OFFSET (page-1)*pageSize
	offset := (page - 1) * pageSize

	// 2. 构建基础 SQL 和参数
	query := "SELECT id, name, age FROM users WHERE 1=1"
	var args []interface{}

	// 3. 动态添加过滤条件
	if nameFilter != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+nameFilter+"%") // 模糊查询
	}

	// 4. 添加分页
	query += " LIMIT ? OFFSET ?"
	args = append(args, pageSize, offset)

	// 5. 执行查询
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 6. 循环处理结果集
	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
