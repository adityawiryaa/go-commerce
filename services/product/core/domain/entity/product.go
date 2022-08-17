package entity

import (
	"go-commerce/pkg/utils"
	"time"
)

type Product struct {
	Id        string
	Name      string
	Quantity  int32
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ProductDTO struct {
	Id        string
	Name      string
	Quantity  int32
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(payload ProductDTO) *Product {
	var (
		id        string    = payload.Id
		createdAt time.Time = payload.CreatedAt
		updatedAt time.Time = payload.UpdatedAt
	)

	if payload.Id == "" {
		id = utils.GenerateUUID()
	}
	if createdAt.IsZero() {
		createdAt = time.Now()
	}
	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	product := &Product{
		Id:        id,
		Name:      payload.Name,
		Quantity:  payload.Quantity,
		Price:     payload.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	return product
}
