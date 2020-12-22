package main

import (
	"encoding/json"
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
	//ecnoインスタンスの生成 -> express的な
	e := echo.New()

	//ミドルウェア関数の登録
	// e.Use(middleware.Logger())
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

	//GET - /api/lists
	e.GET("/api/lists", func(c echo.Context) error {
		users := []List{}
		db.Find(&users)
		res, err := json.Marshal(users)
		if err != nil {
			panic("Error in GET/api/lists")
		}
		return c.JSONBlob(http.StatusOK, res)
	})

	//POST - /api/lists
	e.POST("api/lists", func(c echo.Context) error {
		req := new(List)
		err := c.Bind(&req)
		if err != nil {
			panic("Error in POST/api.lists")
		}
		list := List{Title: req.Title, Description: req.Description}
		db.Create(&list)
		return c.String(http.StatusOK, "ok")
	})

	//GET - /api/items/:listId
	e.GET("/api/items/:listId", func(c echo.Context) error {
		listId := c.Param("listId")
		items := []Item{}
		db.Where("list_id = ?", listId).Find(&items)
		res, err := json.Marshal(items)
		if err != nil {
			panic("Error in GET/api/items/:listId")
		}
		return c.JSONBlob(http.StatusOK, res)
	})

	//POST - /api/items
	e.POST("/api/items", func(c echo.Context) error {
		req := new(Item)
		err := c.Bind(&req)
		if err != nil {
			panic("Error in POST/api/items")
		}
		item := Item{Title: req.Title, Description: req.Description, Url: req.Url, ListID: req.ListID}
		db.Create(&item)
		return c.String(http.StatusOK, "ok")
	})
	//port3000でサーバーをたてる
	e.Logger.Fatal(e.Start(":3000"))
}
