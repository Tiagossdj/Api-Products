package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var prodcutObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&prodcutObj.ID,
			&prodcutObj.Name,
			&prodcutObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, prodcutObj)
	}

	rows.Close()

	return productList, err

}
