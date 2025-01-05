package repository

import (
	"context"
	"database/sql"
	"fmt"

	"wetees.com/domain"
)

const (
	TABLE_PRODUCTS = "products"
)

type productRepository struct {
	Conn *sql.DB
}

// Get products
func (p *productRepository) GetProducts(c context.Context, page int, limit int) ([]domain.Product, int, error) {
	result := []domain.Product{}

	// Calculate all rows count
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", TABLE_PRODUCTS)
	if err := p.Conn.QueryRowContext(c, query).Scan(&count); err != nil {
		return result, count, err
	}

	// Calculate the offset
	offset := (page - 1) * limit
	query = fmt.Sprintf("SELECT id, name, category_id, merchant_id, price, status, created_at FROM %s WHERE status=? LIMIT %d OFFSET %d", TABLE_PRODUCTS, limit, offset)
	// Get only products with status 'available'
	rows, err := p.Conn.QueryContext(c, query, 1)
	if err != nil {
		return result, count, err
	}
	defer rows.Close()
	for rows.Next() {
		product := domain.Product{}
		if err = rows.Scan(&product.ID, &product.Name, &product.CategoryId, &product.MerchantId, &product.Price, &product.Status, &product.CreatedAt); err == nil {
			result = append(result, product)
		}
	}

	return result, count, nil
}

// Search products
func (p *productRepository) SearchProducts(c context.Context, item string, page int, limit int) ([]domain.Product, int, error) {
	result := []domain.Product{}

	// Calculate all rows count
	var count int
	fullNameQuery := "%" + item + "%"
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE status=? AND name LIKE ? ", TABLE_PRODUCTS)
	if err := p.Conn.QueryRowContext(c, query, 1, fullNameQuery).Scan(&count); err != nil {
		return result, count, err
	}

	offset := (page - 1) * limit
	query = fmt.Sprintf("SELECT id, name, category_id, merchant_id, price, status, created_at FROM %s WHERE status=? AND name LIKE ? LIMIT %d OFFSET %d", TABLE_PRODUCTS, limit, offset)
	// Get only products with status 'available'
	rows, err := p.Conn.QueryContext(c, query, 1, fullNameQuery)
	if err != nil {
		return result, count, err
	}
	defer rows.Close()

	for rows.Next() {
		product := domain.Product{}
		if err = rows.Scan(&product.ID, &product.Name, &product.CategoryId, &product.MerchantId, &product.Price, &product.Status, &product.CreatedAt); err == nil {
			result = append(result, product)
		}
	}

	return result, count, nil
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{
		Conn: db,
	}
}
