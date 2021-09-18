package config

import (
	"o-rest/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func DBInit(username, password, address, port, name string) *gorm.DB {
	connectionString := username+":"+password+"@tcp("+address+":"+port+")/"+name+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql",connectionString)

	if err != nil {
		panic("failed to connect db" + err.Error())
	}

	db.AutoMigrate(entity.Item{}, entity.Order{})
	return db
}
