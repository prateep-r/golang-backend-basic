curl --location --request POST 'http://localhost:8080/api/v1/product' \
--header 'Content-Type: application/json' \
--data '{
    "productName": "xxx",
    "price": 123.23
}'