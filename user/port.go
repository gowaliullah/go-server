package user

import (
	"github.com/gowaliullah/ecommerce/domain"
	serviceHandler "github.com/gowaliullah/ecommerce/rest/handlers/user"
)

type Service interface {
	serviceHandler.Servece // embedding
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)

	// Get(userId int) (*User, error)
	// List() ([]*User, error)
	// Delete(userId int) error
	// Update(user User) (*User, error)
}
