package exercise

import "training/pointer_func"

// Ex08 /* เขียน function หาราคารวมของ Product ทั้งหมด
func Ex08(productList []Product) float64 {
	sum := 0.0
	for _, product := range productList {
		sum += pointer_func.ToValue(product.ProductPrice, 0)
	}

	return sum
}
