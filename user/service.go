package user

import "github.com/gowaliullah/ecommerce/domain"

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.usrRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}
func (svc *service) Find(email, password string) (*domain.User, error) {
	usr, err := svc.usrRepo.Find(email, password)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}

	return usr, nil
}
