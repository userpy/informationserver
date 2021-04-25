package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB //база данных

func init() {
	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}
	dbName := os.Getenv("db_name")
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{}, &Role{}) //Миграция базы данных
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}
