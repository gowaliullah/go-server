package user

import "github.com/gowaliullah/basic-ecommerce/domain"

type Servece interface {
	Create(user domain.User) (*domain.User, error)
	Find(Email, Password string) (*domain.User, error)
}
