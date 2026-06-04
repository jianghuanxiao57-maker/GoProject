// 路径：GoProject/Main.go
package main

import (
	controllers "GoProject/Controllers" // 确保这里的大小写和左侧文件夹完全一致
	database "GoProject/db"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("gin_server.log")
	// 设置 Gin 的输出：同时输出到控制台和文件
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	database.InitDB() // 执行初始化

	r := gin.Default()
	// 用户模块接口 (v1版本)
	userGroup := r.Group("/api/v1/user")
	{
		// 这里的路径会拼接成 /api/v1/user/list
		userGroup.GET("/list", controllers.GetUserListHandler)
		// 增加新接口：比如修改用户
		// userGroup.POST("/update", controllers.UpdateUserHandler)
		// // 增加新接口：删除用户
		// userGroup.DELETE("/delete", controllers.DeleteUserHandler)
	}

	// // 订单模块接口
	// orderGroup := r.Group("/api/v1/order")
	// {
	// 	orderGroup.GET("/list", controllers.GetOrderList)
	// 	orderGroup.POST("/create", controllers.CreateOrder)
	// }

	// --- 路由分组结束 ---

	r.Run(":8080")
}
