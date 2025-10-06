package product

import (
	"github.com/gowaliullah/basic-ecommerce/domain"
	pdrHanlar "github.com/gowaliullah/basic-ecommerce/rest/handlers/product"
)

type Service interface {
	pdrHanlar.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Delete(prdId int) error
	Update(p domain.Product) (*domain.Product, error)
}
