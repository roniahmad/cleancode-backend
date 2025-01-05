package domain

import (
	"context"
	"database/sql"
)

type Product struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	CategoryId int          `json:"category_id"`
	MerchantId int          `json:"merchant_id"`
	Price      float64      `json:"price"`
	Status     int          `json:"status"`
	CreatedAt  sql.NullTime `json:"created_at"`
}

type ProductRepository interface {
	GetProducts(c context.Context, page, limit int) ([]Product, int, error)
	SearchProducts(c context.Context, item string, page, limit int) ([]Product, int, error)
}

type ProductUsecase interface {
	GetProducts(c context.Context, page, limit int) ([]Product, int, int, error)
	SearchProducts(c context.Context, item string, page, limit int) ([]Product, int, int, error)
}
