package product

import "github.com/google/uuid"

type SaveProductRequest struct {
	ProductName string  `json:"productName" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type SaveProductPayload struct {
	ProductName string
	Price       float64
}

type SaveProductDto struct {
	ProductId   uuid.UUID
	ProductName string
	Price       float64
}

type UpdateProductRequest struct {
	ProductId   string  `uri:"productId" validate:"required,uuid"`
	ProductName string  `json:"productName" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type UpdateProductPayload struct {
	ProductId   uuid.UUID
	ProductName string
	Price       float64
}

type DeleteProductRequest struct {
	ProductId string `uri:"productId" validate:"required,uuid"`
}

type GetProductRequest struct {
	ProductId string `uri:"productId" validate:"required,uuid"`
}

type GetProductDto struct {
	ProductId   uuid.UUID
	ProductName string
	Price       float64
}

type GetProductResponse struct {
	ProductId   string  `json:"productId"`
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
}
