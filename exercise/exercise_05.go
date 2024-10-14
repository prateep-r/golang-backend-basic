package exercise

import (
	"bytes"
	"encoding/json"
)

// Ex05 /* แปลง map[string]any ทีรับเข้ามา เป็น Product แล้ว return ออกไป
func Ex05(productMap map[string]any) Product {

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(productMap)
	if err != nil {
		panic(err)
	}

	var product Product
	err = json.NewDecoder(bytes.NewBuffer(buffer.Bytes())).Decode(&product)
	if err != nil {
		panic(err)
	}

	return product
}
