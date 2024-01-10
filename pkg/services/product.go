package services

import (
	"github.com/fdrt29/product-app/pkg"
	"github.com/fdrt29/product-app/pkg/repositories"
)

type ProductService struct {
	repo repositories.Producter
}

func NewProductService(repo repositories.Producter) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product pkg.Product) (int, error) {
	return s.repo.Create(product)
}

func (s *ProductService) GetAll() ([]pkg.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetById(productId int) (pkg.Product, error) {
	return s.repo.GetById(productId)
}

func (s *ProductService) Update(productId int, input pkg.UpdateProductInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(productId, input)
}

func (s *ProductService) Delete(productId int) error {
	return s.repo.Delete(productId)
}
