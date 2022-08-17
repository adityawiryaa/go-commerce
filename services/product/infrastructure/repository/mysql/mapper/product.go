package mapper

import (
	"go-commerce/services/product/core/domain/entity"
	"go-commerce/services/product/infrastructure/repository/mysql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToModelProduct(product *entity.Product) *models.Product {
	return &models.Product{
		Id:        product.Id,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		DeletedAt: product.DeletedAt,
	}
}

func ToDomainProduct(models *models.Product) *entity.Product {
	return &entity.Product{
		Id:        models.Id,
		Name:      models.Name,
		Quantity:  models.Quantity,
		Price:     models.Price,
		CreatedAt: models.CreatedAt,
		UpdatedAt: models.UpdatedAt,
		DeletedAt: models.DeletedAt,
	}
}

func ToDomainProductList(models []*models.Product) []*entity.Product {
	var products []*entity.Product
	for _, model := range models {
		products = append(products, ToDomainProduct(model))
	}
	return products
}

func ToDbqStructProduct(product *entity.Product) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqProduct(product))
	return
}

func DataDbqProduct(product *entity.Product) []interface{} {
	return dbq.Struct(ToModelProduct(product))
}
