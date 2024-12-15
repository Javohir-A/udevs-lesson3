package models

type Order struct {
	ID         string           `json:"id" bson:"_id,omitempty"`
	CustomerID string           `json:"customer_id" bson:"customer_id"`
	Products   []ProductInOrder `json:"products" bson:"products"`
	TotalPrice float64          `json:"total_price" bson:"total_price"`
	OrderDate  string           `json:"order_date" bson:"order_date"`
	Status     string           `json:"status" bson:"status"`
	CreatedAt  string           `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  string           `json:"updated_at" bson:"updated_at,omitempty"`
}

type ProductInOrder struct {
	ProductID string  `json:"product_id" bson:"product_id"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float64 `json:"price" bson:"price"`
}
