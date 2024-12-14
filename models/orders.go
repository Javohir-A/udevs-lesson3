package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`                // MongoDB ObjectID for the order
	CustomerID string             `json:"customer_id" bson:"customer_id"`         // ID of the customer placing the order
	Products   []ProductInOrder   `json:"products" bson:"products"`               // List of products in the order
	TotalPrice float64            `json:"total_price" bson:"total_price"`         // Total price of the order
	OrderDate  time.Time          `json:"order_date" bson:"order_date"`           // Timestamp when the order was placed
	Status     string             `json:"status" bson:"status"`                   // Status of the order (e.g., "pending", "completed")
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"` // Timestamp of order creation
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at,omitempty"` // Timestamp of last update
}

type ProductInOrder struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"` // Reference to the Product's ID
	Quantity  int                `json:"quantity" bson:"quantity"`     // Quantity of the product ordered
	Price     float64            `json:"price" bson:"price"`           // Price of the product at the time of order
}
