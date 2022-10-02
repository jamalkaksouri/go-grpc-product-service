package db

import (
	"github.com/jamalkaksouri/go-grpc-product-svc/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	_ = db.AutoMigrate(&models.Product{})
	_ = db.AutoMigrate(&models.StockDecreaseLog{})

	return Handler{db}
}
