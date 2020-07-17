package dao

import (
	"demo_order/db"
	"demo_order/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// 新增
func TestOrderMysqlDao_CreateOrder(t *testing.T) {

	dbLink, _ := db.Init()
	Dao := &OrderMysqlDao{Db: dbLink}

	err := Dao.CreateOrder(&model.Order{
		ID:        0,
		OrderId:   "",
		UserName:  "",
		Amount:    0,
		Status:    "",
		FileUrl:   "",
		CreatedAt: time.Now().Unix(),
	})
	assert.NoError(t, err)
}

// 数据
func TestOrderMysqlDao_QueryOrder(t *testing.T) {

	dbLink, _ := db.Init()
	Dao := &OrderMysqlDao{Db: dbLink}

	id := 10
	order, err := Dao.QueryOrder(int64(id))
	assert.NoError(t, err)
	assert.Equal(t, int64(10), order.ID)
}

// 列表
func TestOrderMysqlDao_QueryOrders(t *testing.T) {

	dbLink, _ := db.Init()
	Dao := &OrderMysqlDao{Db: dbLink}

	username := "xiaoming"
	page := 0
	limit := 0
	_, err := Dao.QueryOrders(username, page, limit)
	assert.NoError(t, err)
}

// 更新
func TestOrderMysqlDao_UpdateOrder(t *testing.T) {

	dbLink, _ := db.Init()
	Dao := &OrderMysqlDao{Db: dbLink}

	order := &model.Order{
		ID:        10,
		OrderId:   "10",
		UserName:  "10",
		Amount:    10,
		Status:    "10",
		FileUrl:   "10",
		CreatedAt: time.Now().Unix(),
	}

	err := Dao.UpdateOrder(order)
	assert.NoError(t, err)

}

// 更新url
func TestOrderMysqlDao_UpdateUrl(t *testing.T) {

	dbLink, _ := db.Init()
	Dao := &OrderMysqlDao{Db: dbLink}

	id := 10
	url := "upload/2020/7/14/1594720514-用户类图.jpg"

	err := Dao.UpdateUrl(int64(id), url)
	assert.NoError(t, err)
}
