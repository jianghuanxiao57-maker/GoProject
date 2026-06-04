// 路径：GoProject/Controllers/user_handler.go
package controllers

import (
	"GoProject/Query"
	"GoProject/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserListHandler(c *gin.Context) {
	// 1. 获取过滤参数 (不再强制要求 id)
	name := c.Query("name")

	// 2. 获取分页参数并转换类型
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 3. 调用你写在 Query 包里的新函数
	users, err := Query.QueryUserList(db.Conn, name, page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询失败"})
		return
	}

	// 4. 返回 JSON 数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": users,
		"pagination": gin.H{
			"page":     page,
			"pageSize": pageSize,
		},
	})
}
