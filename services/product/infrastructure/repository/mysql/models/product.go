package models

import "time"

type Product struct {
	Id        string     `dbq:"id"`
	Name      string     `dbq:"name"`
	Quantity  int32      `dbq:"quantity"`
	Price     int64      `dbq:"price"`
	CreatedAt time.Time  `dbq:"created_at"`
	UpdatedAt time.Time  `dbq:"updated_at"`
	DeletedAt *time.Time `dbq:"deleted_at"`
}

func (Product) TableName() string {
	return "products"
}

func TableProducts() []string {
	return []string{
		"id",
		"name",
		"quantity",
		"price",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
