package tools

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "kinny"
	PASSWORD = "kinny123"
	NETWORK  = "tcp"
	// SERVER   = "192.168.0.186"	//本機
	// SERVER = "35.224.139.168" //public DB
	SERVER   = "10.38.80.9" //private DB
	PORT     = 3306
	DATABASE = "kinny_shop"
)

func Ormconnect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
		//return nil
	}
	return db
}
func Sqlconnect() *sql.DB {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		return nil
	}
	if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為：", err.Error())
		return nil
	}
	return db
}
