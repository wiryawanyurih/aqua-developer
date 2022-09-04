package connection

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectio() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=AquaYurih123 dbname=efishery_academy port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
		return db, err
	}
	db.AutoMigrate(&model.customer{})

	return db, nil

}
