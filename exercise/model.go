package exercise

import "time"

type ProductUnit string

const (
	Can    ProductUnit = "CAN"
	Bottom ProductUnit = "BOTTOM"
	Glass  ProductUnit = "GLASS"
)

type Product struct {
	ProductId    int64       `json:"productId"`
	ProductCode  string      `json:"productCode"`
	ProductName  string      `json:"productName"`
	ProductPrice *float64    `json:"productPrice"`
	ProductUnit  ProductUnit `json:"productUnit"`
	CreatedBy    string      `json:"createdBy"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedBy    *string     `json:"updatedBy"`
	UpdatedAt    *time.Time  `json:"updatedAt"`
}

type Order struct {
	OrderId     int64      `json:"orderId"`
	OrderNo     string     `json:"orderNo"`
	ProductList []Product  `json:"productList"`
	CreatedBy   string     `json:"createdBy"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedBy   *string    `json:"updatedBy"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
