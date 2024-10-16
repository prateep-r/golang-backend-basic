package product

import (
	"context"
	"time"
	"training/persistence"

	"github.com/google/uuid"
)

type Service interface {
	Save(ctx context.Context, payload SaveProductPayload) (*SaveProductDto, error)
	Update(ctx context.Context, payload UpdateProductPayload) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*GetProductDto, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Save(ctx context.Context, payload SaveProductPayload) (*SaveProductDto, error) {
	product, err := s.repository.Insert(ctx, persistence.Product{
		ProductName: payload.ProductName,
		Price:       payload.Price,
		CreatedBy:   "ADMIN",
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &SaveProductDto{
		ProductId:   product.ProductId,
		ProductName: product.ProductName,
		Price:       product.Price,
	}, nil
}

func (s *service) Update(ctx context.Context, payload UpdateProductPayload) error {
	user := "ADMIN"
	now := time.Now()

	return s.repository.Update(ctx, persistence.Product{
		ProductId:   payload.ProductId,
		ProductName: payload.ProductName,
		Price:       payload.Price,
		UpdatedBy:   &user,
		UpdatedAt:   &now,
	})
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) GetById(ctx context.Context, id uuid.UUID) (*GetProductDto, error) {
	product, err := s.repository.SelectById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetProductDto{
		ProductId:   product.ProductId,
		ProductName: product.ProductName,
		Price:       product.Price,
	}, nil
}
