package repositories

import (
	"database/sql"
	"fmt"
	"kasir-online/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin() // use DBTransaction to enable rollback
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)
	nomor := 1

	for _, item := range items {
		// query tabel produk untuk mendapatkan harga, stock dan nama produk
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product id %d not found", item.ProductID)
		}
		if item.Quantity > stock {
			return nil, fmt.Errorf("Stock %s unavailable", productName)
		}
		if err != nil {
			return nil, err
		}

		subTotal := productPrice * item.Quantity
		totalAmount += subTotal

		// update: kurangi stock di tabel produk
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		// masukkan detail produk ke dalam TransactionDetail
		details = append(details, models.TransactionDetail{
			ID:          nomor,
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subTotal,
		})
		nomor++
	}

	// save ke tabel transaction
	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// save ke tabel transaction_details
	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)

		if err != nil {
			return nil, err
		}
	}

	// commit DBTransaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// kembalikan transaksi
	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
