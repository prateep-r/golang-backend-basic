package exercise_test

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
	"training/exercise"
	"training/pointer_func"
)

func TestEx(t *testing.T) {

	t.Run("Ex01", func(t *testing.T) {
		got := exercise.Ex01()
		assert.Equal(t, true, got)
	})

	t.Run("Ex02", func(t *testing.T) {
		products := exercise.Ex02()
		assert.Equal(t, 3, len(products))
		assert.Equal(t, int64(1111), products[0].ProductId)
		assert.Equal(t, "COKE", products[0].ProductCode)
		assert.Equal(t, "Coca cola", products[0].ProductName)
		assert.Equal(t, 15.00, products[0].ProductPrice)
		assert.Equal(t, exercise.Can, products[0].ProductUnit)
		assert.Equal(t, "Josh", products[0].CreatedBy)
		assert.NotEqual(t, nil, products[0].CreatedAt)
		assert.Equal(t, "Sarah", products[0].UpdatedBy)
		assert.NotEqual(t, nil, products[0].UpdatedAt)

		assert.Equal(t, int64(2222), products[1].ProductId)
		assert.Equal(t, "PEPSI", products[1].ProductCode)
		assert.Equal(t, "Pepsi", products[1].ProductName)
		assert.Equal(t, 15.50, products[1].ProductPrice)
		assert.Equal(t, exercise.Bottom, products[1].ProductUnit)
		assert.Equal(t, "John", products[1].CreatedBy)
		assert.NotEqual(t, nil, products[1].CreatedAt)
		assert.Equal(t, nil, products[1].UpdatedBy)
		assert.Equal(t, nil, products[1].UpdatedAt)

		assert.Equal(t, int64(3333), products[2].ProductId)
		assert.Equal(t, "SPRITE", products[2].ProductCode)
		assert.Equal(t, "Sprite", products[2].ProductName)
		assert.Equal(t, nil, products[2].ProductPrice)
		assert.Equal(t, exercise.Glass, products[2].ProductUnit)
		assert.Equal(t, "Peter", products[2].CreatedBy)
		assert.NotEqual(t, nil, products[2].CreatedAt)
		assert.Equal(t, nil, products[2].UpdatedBy)
		assert.Equal(t, nil, products[2].UpdatedAt)
	})

	t.Run("Ex03", func(t *testing.T) {
		orders := exercise.Ex03()
		assert.Equal(t, 1, len(orders))
		assert.Equal(t, int64(12345), orders[0].OrderId)
		assert.Equal(t, "ORD0001", orders[0].OrderNo)
		assert.Equal(t, "Tony", orders[0].CreatedBy)
		assert.NotEqual(t, nil, orders[0].CreatedAt)
		assert.Equal(t, nil, orders[0].UpdatedBy)
		assert.Equal(t, nil, orders[0].UpdatedAt)
		assert.Equal(t, 2, len(orders[0].ProductList))

		assert.Equal(t, int64(1111), orders[0].ProductList[0].ProductId)
		assert.Equal(t, "COKE", orders[0].ProductList[0].ProductCode)
		assert.Equal(t, "Coca cola", orders[0].ProductList[0].ProductName)
		assert.Equal(t, 15.00, orders[0].ProductList[0].ProductPrice)
		assert.Equal(t, exercise.Can, orders[0].ProductList[0].ProductUnit)
		assert.Equal(t, "Josh", orders[0].ProductList[0].CreatedBy)
		assert.NotEqual(t, nil, orders[0].ProductList[0].CreatedAt)
		assert.Equal(t, "Sarah", orders[0].ProductList[0].UpdatedBy)
		assert.NotEqual(t, nil, orders[0].ProductList[0].UpdatedAt)

		assert.Equal(t, int64(2222), orders[0].ProductList[1].ProductId)
		assert.Equal(t, "PEPSI", orders[0].ProductList[1].ProductCode)
		assert.Equal(t, "Pepsi", orders[0].ProductList[1].ProductName)
		assert.Equal(t, 15.50, orders[0].ProductList[1].ProductPrice)
		assert.Equal(t, exercise.Bottom, orders[0].ProductList[1].ProductUnit)
		assert.Equal(t, "John", orders[0].ProductList[1].CreatedBy)
		assert.NotEqual(t, nil, orders[0].ProductList[1].CreatedAt)
		assert.Equal(t, nil, orders[0].ProductList[1].UpdatedBy)
		assert.Equal(t, nil, orders[0].ProductList[1].UpdatedAt)
	})

	t.Run("Ex04", func(t *testing.T) {
		count := rand.Intn(100) + 1
		productList := randomProductList(count)

		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(productList)
		if err != nil {
			t.Fatal(err)
		}

		products := exercise.Ex04(buffer.String())
		assert.Equal(t, count, len(products))
	})

	t.Run("Ex05", func(t *testing.T) {
		productRandom := randomProduct(12345)

		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(productRandom)
		if err != nil {
			t.Fatal(err)
		}

		var expectedProduct exercise.Product
		err = json.NewDecoder(bytes.NewBuffer(buffer.Bytes())).Decode(&expectedProduct)
		if err != nil {
			t.Fatal(err)
		}

		var productMap map[string]any
		err = json.NewDecoder(bytes.NewBuffer(buffer.Bytes())).Decode(&productMap)
		if err != nil {
			t.Fatal(err)
		}

		product := exercise.Ex05(productMap)
		assert.Equal(t, expectedProduct, product)
	})

	t.Run("Ex06", func(t *testing.T) {
		count := rand.Intn(1000) + 1
		productList := randomProductList(count)

		index := rand.Intn(count)
		productCode := productList[index].ProductCode
		got := exercise.Ex06(productList, productCode)

		assert.Equal(t, index, got)
	})

	t.Run("Ex07", func(t *testing.T) {
		count := rand.Intn(1000) + 1

		productCodeToDeleted := make([]string, 0)
		expectedProductList := make([]exercise.Product, 0)

		productList := randomProductList(count)
		for _, product := range productList {
			remove := rand.Intn(4)
			if remove == 0 {
				productCodeToDeleted = append(productCodeToDeleted, product.ProductCode)
			} else {
				expectedProductList = append(expectedProductList, product)
			}
		}

		remainingProductList := exercise.Ex07(productList, productCodeToDeleted)

		assert.Equal(t, len(remainingProductList), len(expectedProductList))
	})

	t.Run("Ex08", func(t *testing.T) {
		count := rand.Intn(1000) + 1
		productList := randomProductList(count)

		sum := 0.0
		for _, product := range productList {
			sum += pointer_func.ToValue(product.ProductPrice, 0)
		}

		got := exercise.Ex08(productList)
		assert.Equal(t, sum, got)
	})

	t.Run("Ex09", func(t *testing.T) {
		count := rand.Intn(1000) + 1
		productList := randomProductList(count)

		expectedProductList := make([]exercise.Product, len(productList))
		copy(expectedProductList, productList)

		sort.Slice(expectedProductList, func(i, j int) bool {
			if expectedProductList[i].ProductName == expectedProductList[j].ProductName {
				return expectedProductList[i].ProductCode < expectedProductList[j].ProductCode
			}
			return expectedProductList[i].ProductName < expectedProductList[j].ProductName
		})

		gotProductList := exercise.Ex09(productList)
		assert.Equal(t, len(expectedProductList), len(gotProductList))
		assert.Equal(t, expectedProductList, gotProductList)
	})
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func randomPrice(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomProductUnit() exercise.ProductUnit {
	i := rand.Intn(3)
	switch i {
	case 0:
		return exercise.Can
	case 1:
		return exercise.Bottom
	case 2:
		return exercise.Glass
	default:
		return exercise.Can
	}
}

func randomProduct(id int64) exercise.Product {
	return exercise.Product{
		ProductId:    id,
		ProductCode:  randomString(6),
		ProductName:  randomString(10),
		ProductPrice: pointer_func.ToPointer(randomPrice(1, 1_000_000)),
		ProductUnit:  randomProductUnit(),
		CreatedBy:    randomString(10),
		CreatedAt:    time.Now(),
		UpdatedBy:    pointer_func.ToPointer(randomString(10)),
		UpdatedAt:    pointer_func.ToPointer(time.Now()),
	}
}

func randomProductList(count int) []exercise.Product {
	productList := make([]exercise.Product, count)
	for i := 0; i < count; i++ {
		productList[i] = randomProduct(int64(i))
	}
	return productList
}
