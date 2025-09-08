package cmd

import (
	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/rest"
)

func Serve() {

	cnf := config.GetConfig()

	rest.Start(cnf)

}
