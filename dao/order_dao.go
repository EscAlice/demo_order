package dao

import (
	"demo_order/model"
	"github.com/jinzhu/gorm"
)

type OrderDao interface {

	// 创建订单
	CreateOrder(order *model.Order) error

	// 查询订单
	QueryOrder(id int64) (*model.Order, error)

	// 查询订单全部列表
	QueryAll() ([]*model.Order, error)

	// 查询订单列表
	QueryOrders(username string, page, limit int) ([]*model.Order, error)

	// 更新数据
	UpdateOrder(order *model.Order) error

	// 更新文件路径
	UpdateUrl(id int64, url string) error
}

type OrderMysqlDao struct {
	Db *gorm.DB
}

// 初始化
func NewOrderMysqlDao(db *gorm.DB) *OrderMysqlDao {
	return &OrderMysqlDao{Db: db}
}

// 添加数据
func (m *OrderMysqlDao) CreateOrder(order *model.Order) error {

	tx := m.Db.Begin()
	result := tx.Create(order)
	if result.Error != nil {
		tx.Rollback()
		err := result.Error
		return err
	}
	tx.Commit()
	return nil
}

// 查询数据
func (m *OrderMysqlDao) QueryOrder(id int64) (*model.Order, error) {

	order := &model.Order{
		ID:        0,
		OrderId:   "",
		UserName:  "",
		Amount:    0,
		Status:    "",
		FileUrl:   "",
		CreatedAt: 0,
	}
	query := m.Db.Where("id = ?", id)
	if err := query.First(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

// 查询订单全部列表
func (m *OrderMysqlDao) QueryAll() ([]*model.Order, error) {

	var orders []*model.Order
	if err := m.Db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// 查询数据列表
func (m *OrderMysqlDao) QueryOrders(username string, page, limit int) ([]*model.Order, error) {

	var orders []*model.Order
	query := m.Db.Where("user_name LIKE ?", username).Order("created_at desc, amount")
	query = query.Offset((page - 1) * limit).Limit(limit)
	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// 更新数据
func (m *OrderMysqlDao) UpdateOrder(order *model.Order) error {

	tx := m.Db.Begin()
	if err := tx.Save(order).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 更新文件路径
func (m *OrderMysqlDao) UpdateUrl(id int64, url string) error {

	order := model.Order{
		ID:        0,
		OrderId:   "",
		UserName:  "",
		Amount:    0,
		Status:    "",
		FileUrl:   "",
		CreatedAt: 0,
	}
	tx := m.Db.Begin()
	if err := tx.Model(&order).Where("id = ?", id).Update("file_url", url).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
