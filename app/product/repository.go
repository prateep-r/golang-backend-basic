package product

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"training/persistence"
)

type Repository interface {
	Insert(ctx context.Context, model persistence.Product) error
	Update(ctx context.Context, model persistence.Product) error
	Delete(ctx context.Context, productId uuid.UUID) error
	SelectById(ctx context.Context, productId uuid.UUID) (*persistence.Product, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(ctx context.Context, product persistence.Product) error {
	_, err := r.db.Exec(ctx,
		"insert into product(product_id, product_name, price, created_by, created_at) values($1, $2, $3, $4, $5)",
		product.ProductId,
		product.ProductName,
		product.Price,
		product.CreatedBy,
		product.CreatedAt,
	)
	return err
}

func (r *repository) Update(ctx context.Context, product persistence.Product) error {
	_, err := r.db.Exec(ctx,
		"update product set product_name = $1, price = $2, updated_by = $3, updated_at = $4 where product_id = $5",
		product.ProductName,
		product.Price,
		product.UpdatedBy,
		product.UpdatedAt,
		product.ProductId,
	)
	return err
}

func (r *repository) Delete(ctx context.Context, productId uuid.UUID) error {
	_, err := r.db.Exec(ctx,
		"delete from product where product_id = $1",
		productId,
	)
	return err
}

func (r *repository) SelectById(ctx context.Context, productId uuid.UUID) (*persistence.Product, error) {
	var product persistence.Product
	if err := r.db.QueryRow(ctx,
		"SELECT product_id, product_name, price, created_by, created_at, updated_by, updated_at FROM product WHERE product_id = $1",
		productId,
	).Scan(&product.ProductId,
		&product.ProductName,
		&product.Price,
		&product.CreatedBy,
		&product.CreatedAt,
		&product.UpdatedBy,
		&product.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &product, nil
}
