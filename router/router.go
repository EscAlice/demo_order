package router

import (
	"demo_order/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	// 获取数据列表
	router.GET("/order/list", handler.GetList)

	// 获取数据
	router.GET("/order/one", handler.GetOne)

	// 新增数据
	router.POST("/order/add", handler.NewOne)

	// 更新数据
	router.POST("/order/update", handler.UpdateOne)

	// 文件上传
	router.POST("/order/upload", handler.Upload)

	// 文件下载
	router.POST("/order/download", handler.Download)

	// 文件导出
	router.POST("/order/export", handler.Export)

	return router
}
