package product

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
	"training/app"
	"training/persistence"
)

type Service interface {
	Save(ctx context.Context, payload SaveProductPayload) (*SaveProductDto, error)
	Update(ctx context.Context, payload UpdateProductPayload) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*GetProductDto, error)
}

type service struct {
	repository Repository
	redis      redis.UniversalClient
}

func NewService(repository Repository, redisClient redis.UniversalClient) Service {
	return &service{
		repository: repository,
		redis:      redisClient,
	}
}

func (s *service) Save(ctx context.Context, payload SaveProductPayload) (*SaveProductDto, error) {
	product := persistence.Product{
		ProductId:   uuid.Must(uuid.NewV7()),
		ProductName: payload.ProductName,
		Price:       payload.Price,
		CreatedBy:   "ADMIN",
		CreatedAt:   time.Now(),
	}
	err := s.repository.Insert(ctx, product)
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

	if err := s.repository.Update(ctx, persistence.Product{
		ProductId:   payload.ProductId,
		ProductName: payload.ProductName,
		Price:       payload.Price,
		UpdatedBy:   &user,
		UpdatedAt:   &now,
	}); err != nil {
		return err
	}

	if err := s.removeCacheFromRedis(ctx, payload.ProductId); err != nil {
		fmt.Println("Error deleting from redis")
	}
	return nil
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	if err := s.removeCacheFromRedis(ctx, id); err != nil {
		fmt.Println("Error deleting from redis")
	}
	return nil
}

func (s *service) GetById(ctx context.Context, id uuid.UUID) (*GetProductDto, error) {
	product, err := s.getByIdFromRedis(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		product, err = s.repository.SelectById(ctx, id)
		if err != nil {
			return nil, err
		}

		if product != nil {
			if err := s.cacheToRedis(ctx, *product); err != nil {
				fmt.Println("Error caching product to redis:", err)
			}
		}
	}

	if product != nil {
		return &GetProductDto{
			ProductId:   product.ProductId,
			ProductName: product.ProductName,
			Price:       product.Price,
		}, nil
	}

	return nil, errors.New("product not found")
}

func (s *service) getByIdFromRedis(ctx context.Context, id uuid.UUID) (*persistence.Product, error) {
	byteResult, err := s.redis.Get(ctx, fmt.Sprintf("%v:%v", app.RedisProductKey, id.String())).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	var product persistence.Product
	if err := json.NewDecoder(bytes.NewReader(byteResult)).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *service) cacheToRedis(ctx context.Context, product persistence.Product) error {
	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(product); err != nil {
		return err
	}

	if err := s.redis.Set(ctx, fmt.Sprintf("%v:%v", app.RedisProductKey, product.ProductId.String()), buffer.String(), 10*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

func (s *service) removeCacheFromRedis(ctx context.Context, id uuid.UUID) error {
	if err := s.redis.Del(ctx, fmt.Sprintf("%v:%v", app.RedisProductKey, id.String())).Err(); err != nil {
		return err
	}
	return nil
}
