package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const DatabaseName string = "demo_order"

// 初始化数据库连接
func Init() (*gorm.DB, error) {

	var Db *gorm.DB
	var err error

	// 连接数据库
	Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/"+DatabaseName+"?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	if err != nil {
		// 如果数据库不存在则创建数据库
		createDatabase(DatabaseName)
	}
	return Db, nil
}

// 创建数据库
func createDatabase(name string) {

	// 连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建数据库
	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}

	// 数据库使用权限
	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
}
