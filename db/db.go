package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 初始化数据库连接
func Init() (*gorm.DB, error) {
	var Db *gorm.DB
	var err error
	Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/demo_order?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		return nil, err
	}

	if Db.Error != nil {
		return nil, err
	}
	return Db, nil
}
