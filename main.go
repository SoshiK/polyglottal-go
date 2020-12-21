package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//ecnoインスタンスの生成 -> express的な
	e := echo.New()

	//ミドルウェア関数の登録
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//環境変数の読み込み
	err := godotenv.Load(".env")
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
	fmt.Println((db))
	//GET - /
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//port3000でサーバーをたてる
	e.Logger.Fatal(e.Start(":3000"))
}
