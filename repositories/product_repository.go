package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-online/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

var products = []models.Product{
	{ID: 1, Name: "Indomie", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, price, stock, category_id FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0) // init array

	// isi array
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := "SELECT id, name, price, stock FROM products WHERE id = $1"
	var p models.Product
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("Produk tidak ditemukan")
	}
	if err != nil { // catch other error
		return nil, err
	}
	return &p, nil
}

func (repo *ProductRepository) Create(newProduct *models.Product) (*int, error) {
	query := "INSERT INTO products ( updated_at, name, price, stock, category_id) VALUES (NOW(), $1, $2, $3, $4) RETURNING id"

	err := repo.db.QueryRow(query, newProduct.Name, newProduct.Price, newProduct.Stock, newProduct.CategoryID).Scan(&newProduct.ID)
	if err != nil {
		return nil, err
	}

	return &newProduct.ID, nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"

	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Product ID:%d not found", id)
	}

	return nil
}

func (repo *ProductRepository) Update(id int, updateProduct *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4, updated_at = NOW() WHERE id = $5"

	result, err := repo.db.Exec(query, updateProduct.Name, updateProduct.Price, updateProduct.Stock, updateProduct.CategoryID, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Product ID:%d not found", id)
	}

	return nil
}

// ============= V1 (no database) ===============
func (repo *ProductRepository) GetAll_V1() ([]models.Product, error) {
	return products, nil
}

func (repo *ProductRepository) GetByID_V1(id int) (*models.Product, error) {
	for _, p := range products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("not found")
}
func (repo *ProductRepository) Create_V1(newProduct *models.Product) error {
	products = append(products, *newProduct)
	return nil
}
func (repo *ProductRepository) Delete_V1(id int) error {
	for i := range products {
		if products[i].ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func (repo *ProductRepository) Update_V1(id int, updateProduct *models.Product) error {
	for i := range products {
		if products[i].ID == id {
			updateProduct.ID = id
			products[i] = *updateProduct
			return nil
		}
	}
	return errors.New("not found")
}
