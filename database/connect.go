package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB{
	dsn := "host=localhost user=postgres password=postgres dbname=shop_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil{
		log.Fatal("Can't connect to database: ", err)
	}
	return db
}

func CloseDB(db *gorm.DB){
	dbInstance, err := db.DB()
	dbInstance.Close()
	if err != nil{
		log.Fatal("Can't close database: ", err)
	}
}