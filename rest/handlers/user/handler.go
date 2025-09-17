package user

import "github.com/gowaliullah/ecommerce/repo"

type Handler struct {
	userRepo repo.UserRepo
}

func NewHandler(userRepo repo.UserRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}
