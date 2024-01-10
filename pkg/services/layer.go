package services

import (
	"github.com/fdrt29/product-app/pkg"
	"github.com/fdrt29/product-app/pkg/repositories"
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

func NewLayer(repos *repositories.Layer) *Layer {
	return &Layer{
		Producter: NewProductService(repos.Producter),
	}
}
