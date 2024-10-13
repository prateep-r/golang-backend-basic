CREATE TABLE product (
    product_id UUID PRIMARY KEY,
    product_name VARCHAR(200),
    price NUMERIC(10,2),
    created_by VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(50),
    updated_at TIMESTAMP
);