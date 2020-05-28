package model

import (
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go-gin-template/config"
)

var db *gorm.DB

func init() {
	dbType := config.GetEnv("DB_TYPE")
	dbName := config.GetEnv("DB_NAME")
	user := config.GetEnv("DB_USER")
	password := config.GetEnv("DB_PASSWORD")
	host := config.GetEnv("DB_HOST")

	var err error
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Fatal(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return strings.Title(defaultTableName)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
