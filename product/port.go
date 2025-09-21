package product

import "github.com/gowaliullah/ecommerce/domain"

type Service interface{}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(prdId int) error
	Update(p domain.Product) (*domain.Product, error)
}
