package main

import (
	"demo_order/dao"
	"demo_order/db"
	"demo_order/handler"
	"demo_order/model"
	"demo_order/router"
	"demo_order/service"
)

func main() {

	// 初始化数据库连接
	dbLink, err := db.Init()
	if err != nil {
		panic("数据库连接错误")
	}
	defer dbLink.Close()
	// 数据库建表
	dbLink.AutoMigrate(&model.Order{})
	// 初始化dao对象与service对象
	mysqlDao := dao.NewOrderMysqlDao(dbLink)
	handler.OService = service.NewOrderService(mysqlDao)

	// 初始化router
	engine := router.InitRouter()
	err = engine.Run(":8000")
	if err != nil {
		println(err)
	}
}
