package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func Landscape() *gorm.DB {
	var (
		dbHost   = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		database = os.Getenv("DB_DATABASE")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, user, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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