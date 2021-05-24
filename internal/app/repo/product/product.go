package product

import (
	"context"
	"database/sql"

	"github.com/harunnryd/tempodoloe/config"
	"github.com/harunnryd/tempodoloe/internal/app/model"
	"gorm.io/gorm"
)

type Product interface {
	GetProductByID(context.Context, int) (model.Product, error)
	DeductProductQuantityByID(context.Context, int, int) (model.Product, error)
}

type product struct {
	cfg      config.Config
	ormPgSQL *gorm.DB
}

func New(cfg config.Config, ormPgSQL *gorm.DB) Product {
	return &product{cfg: cfg, ormPgSQL: ormPgSQL}
}

func (p *product) GetProductByID(ctx context.Context, id int) (productModel model.Product, err error) {
	err = p.ormPgSQL.
		Select("id, name, price, quantity, created_at, updated_at, deleted_at").
		Where("id = ?", id).
		Find(&productModel).
		Error

	return
}

const DeductProductQuantityByIDQuery = `UPDATE products SET quantity = quantity - @quantity WHERE id = @id RETURNING id, name, price, quantity, created_at, updated_at, deleted_at`

func (p *product) DeductProductQuantityByID(ctx context.Context, id int, quantity int) (productModel model.Product, err error) {
	rows, err := p.ormPgSQL.
		WithContext(ctx).
		Raw(DeductProductQuantityByIDQuery, sql.Named("quantity", quantity), sql.Named("id", id)).
		Rows()

	for rows.Next() {
		p.ormPgSQL.ScanRows(rows, &productModel)
	}

	return
}
