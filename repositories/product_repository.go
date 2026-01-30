package repositories

import (
	"database/sql"
	"kasir-online/models"
)

type ProductRepository struct {
	db *sql.DB
}

type e string

func (x e) Error() string { return string(x) }

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

var products = []models.Product{
	{ID: 1, Name: "Indomie", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	return products, nil
}
func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	for _, p := range products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, e("not found")
}
func (repo *ProductRepository) Create(newProduct *models.Product) error {
	products = append(products, *newProduct)
	return nil
}
func (repo *ProductRepository) Delete(id int) error {
	for i := range products {
		if products[i].ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return e("not found")
}

func (repo *ProductRepository) Update(id int, updateProduct *models.Product) error {
	for i := range products {
		if products[i].ID == id {
			updateProduct.ID = id
			products[i] = *updateProduct
			return nil
		}
	}
	return e("not found")
}
