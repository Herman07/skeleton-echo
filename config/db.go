package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func Landscape() *gorm.DB {
	var (
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		database = os.Getenv("DB_DATABASE")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password,port, database )
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
fmt.Println("db connect",db)
	idle, _ := strconv.Atoi(os.Getenv("SET_MAX_IDLE_CONN"))
	open, _ := strconv.Atoi(os.Getenv("SET_MAX_OPEN_CONN"))

	pool, err := db.DB()
	pool.SetMaxIdleConns(idle)
	pool.SetMaxOpenConns(open)

	return db.Debug()
}