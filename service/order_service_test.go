package service

import (
	"demo_order/dao"
	"demo_order/db"
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

	req := ReqAddParam{
		OrderId:  "",
		UserName: "",
		Amount:   "",
		Status:   "",
		FileUrl:  "",
	}

	err := orderService.AddOrder(req)
	assert.NoError(t, err)
}

// 数据
func TestOrderService_GetOrder(t *testing.T) {

	orderService := InitService()

	id := 10
	res, err := orderService.GetOrder(int64(id))
	assert.NoError(t, err)
	assert.Equal(t, int64(10), res.ID)
}

// 列表
func TestOrderService_GetOrders(t *testing.T) {

	orderService := InitService()

	username := "xiaoming"
	page := 0
	limit := 0

	list, err := orderService.GetOrders(username, page, limit)

	assert.NoError(t, err)
	assert.Equal(t, "xiaoming", list[0].UserName)
}

// 更新
func TestOrderService_UpdateOne(t *testing.T) {

	orderService := InitService()

	resp := RespGetParam{
		ID:       11,
		OrderId:  "10",
		UserName: "10",
		Amount:   10,
		Status:   "10",
		FileUrl:  "10",
	}

	err := orderService.UpdateOne(resp)
	assert.NoError(t, err)

}

// 更新url
func TestOrderService_UpdateFileUrl(t *testing.T) {

	orderService := InitService()

	id := 10
	url := "upload/2020/7/14/1594720514-用户类图.jpg"

	err := orderService.UpdateFileUrl(int64(id), url)
	assert.NoError(t, err)
}

// 下载文件
func TestOrderService_DownloadFile(t *testing.T) {

	orderService := InitService()

	id := 9

	_, err := orderService.DownloadFile(int64(id))
	assert.NoError(t, err)
}

// 文件导出
func TestOrderService_ExportOrder(t *testing.T) {

	orderService := InitService()

	err := orderService.ExportOrder()
	assert.NoError(t, err)
}
