package repositories

import (
	"database/sql"
	"fmt"
	"kasir-online/models"
	"strings"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// better performance, remove n+1 query
func (repo *TransactionRepository) CreateTransactionV2(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin() // use DBTransaction to enable rollback
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// kumpulkan semua product IDs
	productIDs := make([]interface{}, 0, len(items))
	for _, item := range items {
		productIDs = append(productIDs, item.ProductID)
	}

	// buat values untuk query
	var values string
	for i := range productIDs {
		if i > 0 {
			values += ", "
		}
		values += fmt.Sprintf("$%d", i+1)
	}

	// query semua products sekaligus (fix n+1)
	query := fmt.Sprintf("SELECT id, name, price, stock FROM products WHERE id IN (%s)", values)
	rows, err := tx.Query(query, productIDs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// gunakan map untuk menyimpan product
	products := make(map[int]struct {
		name  string
		price int
		stock int
	})

	for rows.Next() {
		var id, price, stock int
		var name string

		if err := rows.Scan(&id, &name, &price, &stock); err != nil {
			return nil, err
		}
		products[id] = struct {
			name  string
			price int
			stock int
		}{name, price, stock}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// persiapkan loop items
	totalAmount := 0
	details := make([]models.TransactionDetail, 0)
	nomor := 1

	for _, item := range items {
		product, exists := products[item.ProductID]

		if !exists {
			return nil, fmt.Errorf("Product id %d not found", item.ProductID)
		}

		if item.Quantity > product.stock {
			return nil, fmt.Errorf("Stock %s unavailable", product.name)
		}

		subTotal := product.price * item.Quantity
		totalAmount += subTotal

		// update: kurang stock di tabel produk
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		// masukkan detail produk ke dalam TransactionDetail
		details = append(details, models.TransactionDetail{
			ID:          nomor,
			ProductID:   item.ProductID,
			ProductName: product.name,
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

	// save ke tabel transaction_detail
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

// gpt5.2
func (repo *TransactionRepository) CreateTransactionCHATGPT(items []models.CheckoutItem) (*models.Transaction, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("items is empty")
	}

	// use DBTransaction to enable rollback
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// ---------------------------------------------
	// Flattening object: kumpulkan ID dan Quantity
	// buat placeholder untuk values dalam query
	// ---------------------------------------------
	values := []string{}
	args := []any{}

	for i, item := range items {
		values = append(values, fmt.Sprintf("($%d,$%d)", i*2+1, i*2+2))
		args = append(args, item.ProductID, item.Quantity)
	}

	// ---------------------------------------------
	// Ambil data Produk dan lock db
	// ---------------------------------------------
	query := `
WITH raw_items(product_id, qty) AS (
    SELECT
        product_id::INT,
        qty::INT
    FROM (
        VALUES ` + strings.Join(values, ",") + `
    ) AS v(product_id, qty)
),
items AS (
    SELECT product_id, SUM(qty) AS qty
    FROM raw_items
    GROUP BY product_id
)
SELECT
    p.id,
    p.name,
    p.price,
    p.stock,
    i.qty
FROM products p
JOIN items i ON i.product_id = p.id
FOR UPDATE;
`

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalAmount := 0
	details := []models.TransactionDetail{}
	nomor := 1
	found := 0

	for rows.Next() {
		var (
			productID int
			name      string
			price     int
			stock     int
			qty       int
		)

		if err := rows.Scan(&productID, &name, &price, &stock, &qty); err != nil {
			return nil, err
		}

		// cek stock
		if stock < qty {
			return nil, fmt.Errorf("stock %s unavailable", name)
		}

		subtotal := price * qty
		totalAmount += subtotal

		// masukkan detail produk ke dalam TransactionDetail
		details = append(details, models.TransactionDetail{
			ID:          nomor,
			ProductID:   productID,
			ProductName: name,
			Quantity:    qty,
			Subtotal:    subtotal,
		})
		nomor++
		found++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if found == 0 {
		return nil, fmt.Errorf("no valid products found")
	}

	// ---------------------------------------------
	// save total amount ke tabel transaction
	// ---------------------------------------------
	var transactionID int
	err = tx.QueryRow(
		"INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id",
		totalAmount,
	).Scan(&transactionID)

	if err != nil {
		return nil, err
	}

	// ---------------------------------------------
	// save details ke transaction details
	// ---------------------------------------------
	values = []string{}
	args = []any{transactionID}
	idx := 2

	for _, d := range details {
		values = append(values,
			fmt.Sprintf("($%d,$%d,$%d)", idx, idx+1, idx+2))
		args = append(args, d.ProductID, d.Quantity, d.Subtotal)
		idx += 3
	}

	query = `
WITH items(product_id, qty, subtotal) AS (
    SELECT
        product_id::INT,
        qty::INT,
        subtotal::INT
    FROM (
        VALUES ` + strings.Join(values, ",") + `
    ) AS v(product_id, qty, subtotal)
)
INSERT INTO transaction_details
    (transaction_id, product_id, quantity, subtotal)
SELECT
    $1,
    product_id,
    qty,
    subtotal
FROM items;
`

	if _, err := tx.Exec(query, args...); err != nil {
		return nil, err
	}

	// ---------------------------------------------
	// update stock
	// ---------------------------------------------
	values = []string{}
	args = []any{}

	for i, item := range items {
		values = append(values, fmt.Sprintf("($%d,$%d)", i*2+1, i*2+2))
		args = append(args, item.ProductID, item.Quantity)
	}

	query = `
WITH raw_items(product_id, qty) AS (
    SELECT
        product_id::INT,
        qty::INT
    FROM (
        VALUES ` + strings.Join(values, ",") + `
    ) AS v(product_id, qty)
),
items AS (
    SELECT product_id, SUM(qty) AS qty
    FROM raw_items
    GROUP BY product_id
)
UPDATE products p
SET stock = p.stock - i.qty
FROM items i
WHERE p.id = i.product_id;
`

	if _, err := tx.Exec(query, args...); err != nil {
		return nil, err
	}

	// ---------------------------------------------
	// commit DBTransaction
	// ---------------------------------------------
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

func (repo *TransactionRepository) TodayReport(startDate string, endDate string) (*models.Report, error) {
	// definisikan hari dengan benar!
	today := time.Now().Format("2006-01-02")
	condition := "DATE(created_at) = '" + today + "'"
	outerCondition := "DATE(t.created_at) = '" + today + "'"
	date := "'TODAY (" + today + ")'"
	args := []any{}
	if startDate != "" && endDate != "" {
		condition = "DATE(created_at) BETWEEN $1::date AND $2::date"
		outerCondition = "DATE(t.created_at) BETWEEN $1::date AND $2::date"
		date = "CONCAT($1::text, ' to ', $2::text)"
		args = append(args, startDate, endDate)
	}

	// hitung total transaksi hari ini
	// hitung total revenue hari ini
	// kumpulkan semua transaksi hari ini, left join dengan transaction_detail, left join dengan product
	query := `SELECT ` + date + ` as Date,
	  	(SELECT SUM(total_amount) FROM transactions WHERE ` + condition + `) as total_revenue,
  		(SELECT COUNT(*) FROM transactions WHERE ` + condition + `) as total_transaction,
  		p.name as most_sold_product_name,
  		SUM(td.quantity) as total_product_sold
	FROM transactions t
  		LEFT JOIN transaction_details td on td.transaction_id = t.id
  		LEFT JOIN products p on td.product_id = p.id
	WHERE ` + outerCondition + `
	GROUP BY p.id, p.name
	ORDER BY total_product_sold DESC LIMIT 1`

	// response
	var report models.Report
	err := repo.db.QueryRow(query, args...).Scan(&report.Date, &report.TotalRevenue, &report.TotalTransaction, &report.MostProduct.Nama, &report.MostProduct.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			dateValue := date
			if startDate != "" && endDate != "" {
				dateValue = startDate + " to " + endDate
			}
			return &models.Report{
				Date:             dateValue,
				TotalRevenue:     0,
				TotalTransaction: 0,
				MostProduct: models.MostSold{
					Nama:     "",
					Quantity: 0,
				},
			}, nil
		}
		return nil, err
	}

	return &report, nil
}
