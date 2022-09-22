package database

import (
	"log"

	"github.com/storyofhis/backend/colok-ubi/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host    = "localhost"
	port    = 1512 // port postgresql
	user    = "postgres"
	pass    = "0000"
	db_name = "tes_golang"
)

func Init() *gorm.DB {
	dbURL := "postgres:0000@tcp(127.0.0.1:1512)/tes_golang"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Books{})
	return db
}
