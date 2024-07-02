package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/KurobaneShin/eulabs/db"
	"github.com/KurobaneShin/eulabs/types"
)

func main() {
	fmt.Println("running seeds...")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db := db.Create()

	product := types.Product{
		Title:       "title",
		Description: "desc",
		Price:       3000,
	}

	if err := db.CreateProduct(&product); err != nil {
		log.Fatal(err)
	}
}
