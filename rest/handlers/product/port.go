package product

import "github.com/gowaliullah/basic-ecommerce/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Update(domain.Product) (*domain.Product, error)
	Delete(id int) error
}
