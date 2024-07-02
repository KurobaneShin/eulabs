package db

import (
	"context"

	"github.com/KurobaneShin/eulabs/types"
)

func (db DB) GetProductById(id string) (types.Product, error) {
	var product types.Product

	err := db.NewSelect().
		Model(&product).
		Where("id = ?", id).
		Scan(context.Background())
	return product, err
}

func (db DB) CreateProduct(product *types.Product) error {
	_, err := db.NewInsert().
		Model(product).
		Exec(context.Background())
	return err
}

func (db DB) UpdateProduct(product *types.Product) error {
	_, err := db.NewUpdate().
		Model(product).
		WherePK().
		Exec(context.Background())
	return err
}

func (db DB) DeleteProduct(id string) error {
	_, err := db.NewDelete().
		Model(&types.Product{}).
		Where("id = ?", id).
		Exec(context.Background())
	return err
}
