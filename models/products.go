package models

type Product struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Name      string  `json:"name" bson:"name"`
	Category  string  `json:"category" bson:"category"`
	Price     float64 `json:"price" bson:"price"`
	Stock     int     `json:"stock" bson:"stock"`
	CreatedAt string  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at,omitempty"`
}
