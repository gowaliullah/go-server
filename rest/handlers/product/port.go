package product

import "github.com/gowaliullah/basic-ecommerce/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(domain.Product) (*domain.Product, error)
	Delete(id int) error
}
