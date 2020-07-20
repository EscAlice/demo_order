package service

import (
	"demo_order/dao"
	"demo_order/db"
	"demo_order/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 初始化数据库连接
func InitService() *OrderService {

	dbLink, _ := db.Init()
	Dao := &dao.OrderMysqlDao{Db: dbLink}
	orderService := &OrderService{dao: Dao}
	return orderService
}

// 新增
func TestOrderService_AddOrder(t *testing.T) {

	orderService := InitService()

	req := model.AddOrderReq{
		UserName: "",
		Amount:   0,
		Status:   "",
		FileUrl:  "",
	}

	err := orderService.AddOrder(req)
	assert.NoError(t, err)
}

// 数据
func TestOrderService_OrderDetail(t *testing.T) {

	orderService := InitService()

	id := 10
	res, err := orderService.OrderDetail(uint(id))
	assert.NoError(t, err)
	assert.Equal(t, uint(10), res.ID)
}

// 列表
func TestOrderService_OrderList(t *testing.T) {

	orderService := InitService()

	username := "xiaoming"
	page := 0
	limit := 0

	list, err := orderService.OrderList(username, page, limit)

	assert.NoError(t, err)
	assert.Equal(t, "xiaoming", list[0].UserName)
}

// 更新
func TestOrderService_UpdateOrder(t *testing.T) {

	orderService := InitService()

	req := OrderReq{
		ID:       11,
		OrderId:  "10",
		UserName: "10",
		Amount:   10,
		Status:   "10",
		FileUrl:  "10",
	}

	err := orderService.UpdateOrder(req)
	assert.NoError(t, err)

}

// 更新url
func TestOrderService_UpdateFileUrl(t *testing.T) {

	orderService := InitService()

	id := 10
	url := "upload/2020/7/14/1594720514-用户类图.jpg"

	err := orderService.UpdateFileUrl(uint(id), url)
	assert.NoError(t, err)
}

// 下载文件
func TestOrderService_DownloadFile(t *testing.T) {

	orderService := InitService()

	id := 9

	_, err := orderService.DownloadFile(uint(id))
	assert.NoError(t, err)
}

// 文件导出
func TestOrderService_ExportOrder(t *testing.T) {

	orderService := InitService()

	err := orderService.ExportOrder()
	assert.NoError(t, err)
}
