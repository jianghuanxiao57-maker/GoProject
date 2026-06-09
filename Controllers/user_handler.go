package controllers

import (
	"GoProject/Query"
	"GoProject/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserListHandler 处理用户列表请求
func GetUserListHandler(c *gin.Context) {
	// 1. 获取过滤参数 (可选)
	nameFilter := c.Query("name")

	// 2. 获取分页参数并给默认值
	// 如果前端没传 page，默认第 1 页；没传 pageSize，默认每页 10 条
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 3. 调用你刚才写好的 QueryUserList 函数
	users, err := Query.QueryUserList(db.Conn, nameFilter, page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"msg":   "数据库查询失败",
			"error": err.Error(),
		})
		return
	}

	// 4. 返回列表数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": users, // 这里返回的是数组 []User
		"meta": gin.H{
			"page":     page,
			"pageSize": pageSize,
		},
	})
}
