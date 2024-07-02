package main

import (
	"log"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"

	"github.com/KurobaneShin/eulabs/db"
	"github.com/KurobaneShin/eulabs/handlers"
)

func main() {
	db := db.Create()
	productHandler := handlers.NewProductHandler(db)

	e := echo.New()
	e.GET("/product/:id", handlers.Make(productHandler.HandleGetProduct))
	e.POST("/product", handlers.Make(productHandler.HandleCreateProduct))
	e.PUT("/product/:id", handlers.Make(productHandler.HandleUpdateProduct))
	e.DELETE("/product/:id", handlers.Make(productHandler.HandleDeleteProduct))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
