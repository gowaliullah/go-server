package user

import (
	"github.com/gowaliullah/basic-ecommerce/config"
)

type Handler struct {
	cnf *config.Config
	svc Servece
}

func NewHandler(cnf *config.Config, svc Servece) *Handler {
	return &Handler{
		cnf: cnf,
		svc: svc,
	}
}
