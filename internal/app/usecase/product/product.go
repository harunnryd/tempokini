package product

import (
	"context"

	"github.com/harunnryd/tempodoloe/internal/app/model"
	"github.com/harunnryd/tempodoloe/internal/app/repo"
)

type Product interface {
	GetProductByID(context.Context, int) (model.Product, error)
	DeductProductQuantityByID(context.Context, int, int) (model.Product, error)
}

type product struct {
	repo repo.Repo
}

func New(repo repo.Repo) Product {
	return &product{repo: repo}
}

func (p *product) GetProductByID(ctx context.Context, id int) (productModel model.Product, err error) {
	productModel, err = p.repo.Product().GetProductByID(ctx, id)

	return
}

const DeductProductQuantityByIDQuery = `UPDATE products SET quantity = quantity - @quantity WHERE id = @id LIMIT 1 RETURNING id, name, price, quantity, created_at, updated_at, deleted_at`

func (p *product) DeductProductQuantityByID(ctx context.Context, id int, quantity int) (productModel model.Product, err error) {
	productModel, err = p.repo.Product().DeductProductQuantityByID(ctx, id, quantity)

	return
}
