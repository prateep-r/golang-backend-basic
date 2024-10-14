package exercise

import (
	"time"
	"training/pointer_func"
)

// Ex02 /* ให้ return slice of product ตามโครงสร้าง json ดังนี้
//
//	[{
//		"productId": 1111,
//		"productCode": "COKE",
//		"productName": "Coca cola",
//		"productPrice": 15.00,
//		"productUnit": "CAN",
//		"createdBy": "Josh",
//		"createdAt": <today>,
//		"updatedBy": "Sarah",
//		"updatedAt": <today>
//	}, {
//		"productId": 2222,
//		"productCode": "PEPSI",
//		"productName": "Pepsi",
//		"productPrice": 15.50,
//		"productUnit": "BOTTOM",
//		"createdBy": "John",
//		"createdAt": <today>,
//		"updatedBy": null,
//		"updatedAt": null
//	}, {
//		"productId": 3333,
//		"productCode": "SPRITE",
//		"productName": "Sprite",
//		"productPrice": null,
//		"productUnit": "GLASS",
//		"createdBy": "Peter",
//		"createdAt": <today>,
//		"updatedBy": null,
//		"updatedAt": null
//	}]
//
// */
func Ex02() []Product {
	return []Product{
		{
			ProductId:    1111,
			ProductCode:  "COKE",
			ProductName:  "Coca cola",
			ProductPrice: pointer_func.ToPointer(15.00),
			ProductUnit:  Can,
			CreatedBy:    "Josh",
			CreatedAt:    time.Now(),
			UpdatedBy:    pointer_func.ToPointer("Sarah"),
			UpdatedAt:    pointer_func.ToPointer(time.Now()),
		},
		{
			ProductId:    2222,
			ProductCode:  "PEPSI",
			ProductName:  "Pepsi",
			ProductPrice: pointer_func.ToPointer(15.50),
			ProductUnit:  Bottom,
			CreatedBy:    "John",
			CreatedAt:    time.Now(),
			UpdatedBy:    nil,
			UpdatedAt:    nil,
		},
		{
			ProductId:    3333,
			ProductCode:  "SPRITE",
			ProductName:  "Sprite",
			ProductPrice: nil,
			ProductUnit:  Glass,
			CreatedBy:    "Peter",
			CreatedAt:    time.Now(),
			UpdatedBy:    nil,
			UpdatedAt:    nil,
		},
	}
}
