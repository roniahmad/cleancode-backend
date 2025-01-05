package usecase

import (
	"context"
	"time"

	"wetees.com/domain"
)

type productUsecase struct {
	repo           domain.ProductRepository
	conf           *domain.Config
	contextTimeout time.Duration
}

// Get products
func (uc *productUsecase) GetProducts(c context.Context, page int, limit int) ([]domain.Product, int, int, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	totalPages := 0
	products, count, err := uc.repo.GetProducts(ctx, page, limit)
	if err != nil {
		return products, count, totalPages, err
	}
	totalPages = (count + limit - 1) / limit

	return products, count, totalPages, nil
}

// SearchProducts implements domain.ProductUsecase.
func (uc *productUsecase) SearchProducts(c context.Context, item string, page int, limit int) ([]domain.Product, int, int, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	totalPages := 0
	products, count, err := uc.repo.SearchProducts(ctx, item, page, limit)
	if err != nil {
		return products, count, totalPages, err
	}
	totalPages = (count + limit - 1) / limit

	return products, count, totalPages, nil
}

func NewProductUsecase(repo domain.ProductRepository, conf *domain.Config, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		repo:           repo,
		conf:           conf,
		contextTimeout: timeout,
	}
}
