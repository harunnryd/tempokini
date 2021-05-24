package repo

import (
	"log"

	"github.com/harunnryd/tempodoloe/config"
	"github.com/harunnryd/tempodoloe/internal/app/driver/db"
	"github.com/harunnryd/tempodoloe/internal/app/repo/product"
)

type Repo interface {
	Product() product.Product
}

type repo struct {
	product product.Product
}

func NewRepo(cfg config.Config) Repo {
	dbase := db.New(db.WithConfig(cfg))
	pgsqlConn, err := dbase.Manager(db.PgsqlDialectParam)

	if err != nil {
		log.Fatalln("error1", err)
	}

	repo := new(repo)
	repo.product = product.New(cfg, pgsqlConn)

	return repo
}

func (r *repo) Product() product.Product {
	return r.product
}
