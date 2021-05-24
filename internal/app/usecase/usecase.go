package usecase

import (
	"github.com/harunnryd/tempodoloe/internal/app/repo"
	"github.com/harunnryd/tempodoloe/internal/app/usecase/product"
)

type Usecase interface {
	Product() product.Product
}

type usecase struct {
	product product.Product
}

func NewUsecase(repo repo.Repo) Usecase {
	u := new(usecase)

	u.product = product.New(repo)

	return u
}

func (u *usecase) Product() product.Product {
	return u.product
}
