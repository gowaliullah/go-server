package product

import "github.com/gowaliullah/basic-ecommerce/domain"

type service struct {
	prductRepo ProductRepo
}

func NewService(prductRepo ProductRepo) Service {
	return &service{
		prductRepo: prductRepo,
	}
}

func (svc *service) Create(prd domain.Product) (*domain.Product, error) {
	return svc.prductRepo.Create(prd)
}
func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.prductRepo.Get(id)
}
func (svc *service) List() ([]*domain.Product, error) {
	return svc.prductRepo.List()
}
func (svc *service) Update(prd domain.Product) (*domain.Product, error) {
	return svc.prductRepo.Update(prd)
}
func (svc *service) Delete(id int) error {
	return svc.prductRepo.Delete(id)
}
