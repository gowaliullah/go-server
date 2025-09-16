package middleware

import "github.com/gowaliullah/ecommerce/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}
}
