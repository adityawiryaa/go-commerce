package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-commerce/services/product/core/domain/entity"
	"go-commerce/services/product/core/domain/port"
	"go-commerce/services/product/infrastructure/repository/mysql/mapper"
	"go-commerce/services/product/infrastructure/repository/mysql/models"

	queryBuilder "github.com/dhianalyusi/dynamic-query-builder"
	"github.com/rocketlaunchr/dbq/v2"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) productRepository {
	return productRepository{db: db}
}

func StoreProduct(ctx context.Context, E dbq.EFn, payload *entity.Product) error {
	productDbq := mapper.ToDbqStructProduct(payload)
	stmt := dbq.INSERTStmt(models.Product{}.TableName(), models.TableProducts(), len(productDbq), dbq.MySQL)
	_, err := E(ctx, stmt, nil, productDbq)
	if err != nil {
		return err
	}
	return nil
}

func (productRepo productRepository) AddProduct(ctx context.Context, product *entity.Product) error {
	var err error
	/* Store Product If Something Happer will rollback */
	_ = dbq.Tx(ctx, productRepo.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		err = StoreProduct(ctx, E, product)
		if err != nil {
			return
		}
		_ = txCommit()
	})
	if err != nil {
		return err
	}
	return nil
}

func (productRepo productRepository) FindProduct(ctx context.Context, payload port.FindOptions) (*entity.Product, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s %s`, models.Product{}.TableName(), BuildOptions(payload))
	opts := &dbq.Options{
		ConcreteStruct: models.Product{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
		SingleResult:   true,
	}
	result := dbq.MustQ(ctx, productRepo.db, stmt, opts)

	if result != nil {
		product := mapper.ToDomainProduct(result.(*models.Product))
		return product, nil
	} else {
		return nil, errors.New("product not found")
	}
}

func (productRepo productRepository) FindProducts(ctx context.Context, payload port.FindOptions) ([]*entity.Product, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s %s`, models.Product{}.TableName(), BuildOptions(payload))
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: models.Product{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, productRepo.db, stmt, opts)
	if result == nil {
		return nil, errors.New("products not found")
	}
	products := mapper.ToDomainProductList(result.([]*models.Product))
	return products, nil
}

func (productRepo productRepository) CountProducts(ctx context.Context, payload port.FindOptions) (uint32, error) {
	var sql string = fmt.Sprintf(`SELECT COUNT(*) FROM %s %s`, models.Product{}.TableName(), BuildOptions(payload))
	var total uint32 = 0
	productRepo.db.QueryRow(sql).Scan(&total)
	return total, nil
}

func BuildOptions(payload port.FindOptions) string {
	var dqb queryBuilder.DQB

	query := dqb.Where(dqb.And(
		dqb.NewExpression("1", "=", "1"), // FOR HANDLING WHERE CONDITION
		dqb.NewExpression("id", "=", payload.Id),
		dqb.NewExpression("name", "=", payload.Name),
	))
	if payload.IsPaginate {
		if payload.Offset <= 0 {
			payload.Offset = 1
		}
		if payload.Limit == 0 {
			payload.Limit = 10
		}
		offset := (payload.Offset - 1) * payload.Limit
		query = query.Limit(payload.Limit).Offset(offset)
	}

	return query.ToString()
}
