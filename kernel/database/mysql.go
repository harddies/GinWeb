package database

import "C"
import (
	"GinWeb/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type mysql struct {
	DB *gorm.DB
}

var MySQL mysql

func init() {

	var err error
	MySQL.DB, err = gorm.Open(config.C.Mysql.DB.Datasource)
	MySQL.DB.LogMode(true)
	MySQL.DB.DB().SetMaxOpenConns(2000)
	MySQL.DB.DB().SetMaxIdleConns(1000)
	if err != nil {
		panic(err)
	}
}
