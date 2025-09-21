package product

import (
	"github.com/gowaliullah/ecommerce/domain"
	pdrHanlar "github.com/gowaliullah/ecommerce/rest/handlers/product"
)

type Service interface {
	pdrHanlar.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(prdId int) error
	Update(p domain.Product) (*domain.Product, error)
}
