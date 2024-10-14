package exercise

import (
	"bytes"
	"encoding/json"
)

// Ex04 /* แปลง json string ทีรับเข้ามา เป็น slice of Product แล้ว return ออกไป
func Ex04(jsonString string) []Product {

	var products []Product
	err := json.NewDecoder(bytes.NewBufferString(jsonString)).Decode(&products)
	if err != nil {
		panic(err)
	}

	return products
}
