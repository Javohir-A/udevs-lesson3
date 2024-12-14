package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`                // MongoDB ObjectID
	Name      string             `json:"name" bson:"name"`                       // Product name
	Category  string             `json:"category" bson:"category"`               // Product category
	Price     float64            `json:"price" bson:"price"`                     // Product price
	Stock     int                `json:"stock" bson:"stock"`                     // Available stock quantity
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at,omitempty"` // Timestamp of creation
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"` // Timestamp of last update
}
