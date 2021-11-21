package common

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/techoc/ginessential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	username := viper.GetString("dataSource.username")
	host := viper.GetString("dataSource.host")
	port := viper.GetString("dataSource.port")
	database := viper.GetString("dataSource.database")
	password := viper.GetString("dataSource.password")
	charset := viper.GetString("dataSource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
