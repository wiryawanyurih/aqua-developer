package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	var dsn = "host=localhost user=postgres password=AquaYurih123 dbname=efishery_academy port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&product.Product{}, &category.Category{})

	return db, nil
}
