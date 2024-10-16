package product

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"training/app"
	"training/persistence"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	Insert(ctx context.Context, model persistence.Product) (*persistence.Product, error)
	Update(ctx context.Context, model persistence.Product) error
	Delete(ctx context.Context, productId uuid.UUID) error
	SelectById(ctx context.Context, productId uuid.UUID) (*persistence.Product, error)
}

type repository struct {
	db    *pgxpool.Pool
	redis redis.UniversalClient
}

func NewRepository(db *pgxpool.Pool, redis redis.UniversalClient) Repository {
	return &repository{
		db:    db,
		redis: redis,
	}
}

func (r *repository) Insert(ctx context.Context, product persistence.Product) (*persistence.Product, error) {
	product.ProductId = uuid.Must(uuid.NewV7())

	_, err := r.db.Exec(ctx,
		"insert into products(product_id, product_name, price, created_by, created_at) values($1, $2, $3, $4, $5)",
		product.ProductId,
		product.ProductName,
		product.Price,
		product.CreatedBy,
		product.CreatedAt,
	)
	return &product, err
}

func (r *repository) Update(ctx context.Context, product persistence.Product) error {
	cmd, err := r.db.Exec(ctx,
		"update products set product_name = $1, price = $2, updated_by = $3, updated_at = $4 where product_id = $5",
		product.ProductName,
		product.Price,
		product.UpdatedBy,
		product.UpdatedAt,
		product.ProductId,
	)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return app.ErrNotFound
	}

	err = r.removeCacheFromRedis(ctx, product.ProductId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, productId uuid.UUID) error {
	cmd, err := r.db.Exec(ctx,
		"delete from products where product_id = $1",
		productId,
	)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return app.ErrNotFound
	}

	err = r.removeCacheFromRedis(ctx, productId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SelectById(ctx context.Context, productId uuid.UUID) (*persistence.Product, error) {

	product, err := r.getByIdFromRedis(ctx, productId)
	if err != nil {
		return nil, err
	}

	if product == nil {
		product, err = r.selectByIdFromDB(ctx, productId)
		if err != nil {
			return nil, err
		}

		if product != nil {
			err = r.cacheToRedis(ctx, *product)
			if err != nil {
				return nil, err
			}
		}
	}
	return product, nil
}

func (r *repository) selectByIdFromDB(ctx context.Context, productId uuid.UUID) (*persistence.Product, error) {
	var product persistence.Product
	if err := r.db.QueryRow(ctx,
		"SELECT product_id, product_name, price, created_by, created_at, updated_by, updated_at FROM products WHERE product_id = $1",
		productId,
	).Scan(&product.ProductId,
		&product.ProductName,
		&product.Price,
		&product.CreatedBy,
		&product.CreatedAt,
		&product.UpdatedBy,
		&product.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, app.ErrNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *repository) getByIdFromRedis(ctx context.Context, id uuid.UUID) (*persistence.Product, error) {
	byteResult, err := r.redis.Get(ctx, fmt.Sprintf("%v:%v", app.RedisProductKey, id.String())).Bytes()
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

func (r *repository) cacheToRedis(ctx context.Context, product persistence.Product) error {
	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(product); err != nil {
		return err
	}

	if err := r.redis.Set(ctx, fmt.Sprintf("%v:%v", app.RedisProductKey, product.ProductId.String()), buffer.String(), 10*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

func (r *repository) removeCacheFromRedis(ctx context.Context, id uuid.UUID) error {
	if err := r.redis.Del(ctx, fmt.Sprintf("%v:%v", app.RedisProductKey, id.String())).Err(); err != nil {
		return err
	}
	return nil
}
