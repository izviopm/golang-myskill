package db

import (
	"book_inventory/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal("Error load env")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Books{})

	data := models.Books{}
	if db.Find(&data).RowsAffected == 0 {
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Books{
		{
			Title:       "Bumi",
			Author:      "Tere Liye",
			Description: "Buku Fantasy karangan tere liye serial Bumi",
			Stock:       10,
		},
		{
			Title:       "Bulan",
			Author:      "Tere Liye",
			Description: "Buku Fantasy karangan tere liye serial Bumi",
			Stock:       12,
		},
		{
			Title:       "Matahari",
			Author:      "Tere Liye",
			Description: "Buku Fantasy karangan tere liye serial Bumi",
			Stock:       7,
		},
		{
			Title:       "Filosofi Teras",
			Author:      "Henry Manampiring",
			Description: "Buku Self Improvement tentan Filsafat Stoa",
			Stock:       8,
		},
		{
			Title:       "Girls in The Dark",
			Author:      "Akiyoshi Rikako",
			Description: "Buku J-Lit genre thriller karya Akiyoshi Rikako",
			Stock:       5,
		},
	}

	for _, book := range data {
		db.Create(&book)
	}
}
