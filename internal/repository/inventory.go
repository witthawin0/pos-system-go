package repository

import (
	"database/sql"
	"errors"

	"github.com/witthawin0/pos-system-go/internal/domain"
)

// ProductRepositoryImpl implements the ProductRepository interface.
type productRepositoryImpl struct {
	db *sql.DB
}

// NewProductRepositoryImpl creates a new instance of ProductRepositoryImpl.
func NewProductRepositoryImpl(db *sql.DB) *productRepositoryImpl {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) GetProductByID(id int) (domain.Product, error) {
	var product domain.Product

	// Prepare the SQL query
	query := "SELECT id, name, price, quantity FROM products WHERE id = ?"

	// Execute the query
	row := r.db.QueryRow(query, id)

	// Scan the result into a Product struct
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where product with given ID is not found
			return domain.Product{}, errors.New("product not found")
		}
		// Handle other errors
		return domain.Product{}, err
	}

	return product, nil
}

func (r *productRepositoryImpl) GetProducts() ([]domain.Product, error) {
	var products []domain.Product

	// Prepare the SQL query
	query := "SELECT id, name, price, quantity FROM products"

	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set and scan each row into a Product struct
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepositoryImpl) SaveProduct(product domain.Product) error {
	// Prepare the SQL query
	query := "INSERT INTO products (name, price, quantity) VALUES (?, ?, ?)"

	// Execute the query with the product data
	_, err := r.db.Exec(query, product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepositoryImpl) UpdateProduct(product domain.Product) error {
	// Prepare the SQL query
	query := "UPDATE products SET name = ?, price = ?, quantity = ? WHERE id = ?"

	// Execute the query with the product data
	_, err := r.db.Exec(query, product.Name, product.Price, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepositoryImpl) DeleteProduct(id int) error {
	// Prepare the SQL query
	query := "DELETE FROM products WHERE id = ?"

	// Execute the query with the product ID
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
