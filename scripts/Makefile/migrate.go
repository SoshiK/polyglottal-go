package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title       string
	Description string
	Items       []Item
}

type Item struct {
	gorm.Model
	Title       string
	Description string
	Url         string
	ListID      uint
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	//DBとの接続
	url := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(url)
	if err != nil {
		panic(err.Error())
	}
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Error connecting DATABASE")
	}
	db.AutoMigrate(&List{}, &Item{})

	db.Create(&List{Title: "test", Description: "test for test"})
	db.Create(&List{Title: "test2", Description: "test2 for test2"})

	db.Create(&Item{Title: "item1", Description: "item1item1", Url: "", ListID: 1})
	db.Create(&Item{Title: "item2", Description: "item2item2", Url: "", ListID: 1})
	db.Create(&Item{Title: "item3", Description: "item3item3", Url: "", ListID: 2})
}
