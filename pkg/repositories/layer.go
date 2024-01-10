package repositories

import (
	"github.com/fdrt29/product-app/pkg"
	"github.com/jmoiron/sqlx"
)

type Producter interface {
	Create(product pkg.Product) (int, error)
	GetAll() ([]pkg.Product, error)
	GetById(productId int) (pkg.Product, error)
	Update(productId int, input pkg.UpdateProductInput) error
	Delete(productId int) error
}

type Layer struct {
	Producter
}

func NewLayer(db *sqlx.DB) *Layer {
	return &Layer{
		Producter: NewProductPostgres(db),
	}
}
