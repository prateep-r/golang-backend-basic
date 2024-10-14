package exercise

import (
	"time"
	"training/pointer_func"
)

// Ex03 /* ให้ return slice of order ตามโครงสร้าง json ดังนี้
//
//	[{
//		"orderId": 12345,
//		"orderNo": "ORD0001",
//		"productList": [
//			{
//				"productId": 1111,
//				"productCode": "COKE",
//				"productName": "Coca cola",
//				"productPrice": 15.00,
//				"productUnit": "CAN",
//				"createdBy": "Josh",
//				"createdAt": <today>,
//				"updatedBy": "Sarah",
//				"updatedAt": <today>
//			}, {
//				"productId": 2222,
//				"productCode": "PEPSI",
//				"productName": "Pepsi",
//				"productPrice": 15.50,
//				"productUnit": "BOTTOM",
//				"createdBy": "John",
//				"createdAt": <today>,
//				"updatedBy": null,
//				"updatedAt": null
//			}
//		],
//		"createdBy": "Tony",
//		"createdAt": <today>,
//		"updatedBy": null,
//		"updatedAt": null
//	}]
//
// */
func Ex03() []Order {
	return []Order{
		{
			OrderId:   12345,
			OrderNo:   "ORD0001",
			CreatedBy: "Tony",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
			UpdatedBy: nil,
			ProductList: []Product{
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
			},
		},
	}
}
