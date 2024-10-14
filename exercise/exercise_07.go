package exercise

// Ex07 /* เขียนฟังก์ชัน ลบ product จาก slice ของ productCodeToDeleted ออกจาก productList ที่มีอยู่ แล้ว return slice of Product ที่่ไมถูกลบออก
func Ex07(productList []Product, productCodeToDeleted []string) []Product {

	deletedMap := make(map[string]bool)
	for _, productCode := range productCodeToDeleted {
		deletedMap[productCode] = true
	}

	remainingProducts := make([]Product, 0)
	for _, product := range productList {
		if !deletedMap[product.ProductCode] {
			remainingProducts = append(remainingProducts, product)
		}
	}

	return remainingProducts
}
