package product

import "github.com/gowaliullah/ecommerce/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	GetProduct(id int) (*domain.Product, error)
	GetProducts() ([]*domain.Product, error)
	UpdateProduct(domain.Product) (*domain.Product, error)
	Delete(id int) error
}
