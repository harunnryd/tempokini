package handler

import (
	"github.com/harunnryd/tempodoloe/config"
	"github.com/harunnryd/tempodoloe/internal/app/handler/product"
	"github.com/harunnryd/tempodoloe/internal/app/usecase"
)

type Handler interface {
	Product() product.Product
}

type handler struct {
	product product.Product
}

func NewHandler(cfg config.Config, usecase usecase.Usecase) Handler {
	h := new(handler)

	h.product = product.New(usecase)

	return h
}

func (h *handler) Product() product.Product {
	return h.product
}
