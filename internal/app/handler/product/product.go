package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/harunnryd/tempodoloe/internal/app/usecase"
)

type Product interface {
	GetProductByID(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
	DeductProductQuantityByID(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
}

type product struct {
	usecase usecase.Usecase
}

func New(usecase usecase.Usecase) Product {
	return &product{usecase: usecase}
}

func (u *product) GetProductByID(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	resp, err = u.usecase.Product().GetProductByID(r.Context(), productID)

	return
}

type DeductProductQuantityByID struct {
	Quantity int `json:"quantity"`
}

func (u *product) DeductProductQuantityByID(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var deductProductQuantityByID DeductProductQuantityByID
	if err = json.NewDecoder(r.Body).Decode(&deductProductQuantityByID); err != nil {
		return
	}

	resp, err = u.usecase.Product().DeductProductQuantityByID(r.Context(), productID, deductProductQuantityByID.Quantity)

	return
}
