package cmd

import (
	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/rest"
	"github.com/gowaliullah/ecommerce/rest/handlers/product"
	"github.com/gowaliullah/ecommerce/rest/handlers/user"
)

func Serve() {

	cnf := config.GetConfig()
	productHandler := product.NewHaneder()
	userHandler := user.NewFunc()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()

}

// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2lkZW50aXR5dG9vbGtpdC5nb29nbGVhcGlzLmNvbS9nb29nbGUuaWRlbnRpdHkuaWRlbnRpdHl0b29sa2l0LnYxLklkZW50aXR5VG9vbGtpdCIsImlhdCI6MTc1NzM5NTUxMywiZXhwIjoxNzU3Mzk5MTEzLCJpc3MiOiJmaXJlYmFzZS1hZG1pbnNkay02cjM0eUB0YWJuaW5lLWF1dGgtMzQwMDE1LmlhbS5nc2VydmljZWFjY291bnQuY29tIiwic3ViIjoiZmlyZWJhc2UtYWRtaW5zZGstNnIzNHlAdGFibmluZS1hdXRoLTM0MDAxNS5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsInVpZCI6IlowM1NlTXpXSEdaVEVXNGtPaFRKUFlkT0NmdzEifQ.r9dLQU3gOmvvkKwFYvrdXEO-o5levCXIi7ux4D-mIthEJmXZx0zsqCviL_CZiGsl0_pvrRC-b-4jUC62FIp1HSaQoHR1BOH6dd47m3BCgLicaHenBTP6IynCZV7Wl5fjfKyt5qSOLZOTrtop0ZEcysy7j7GjFILzj_Z5iPrIGSUhDiJhH-lImBMc7E-V9tePNCC9aWRNZFWqu1hhIce-AKRivobsrSm10CkMdjBOLk9IXnGzfp-yK6Q5ayZYIdNadkaG96nhmBkokZdMp8SSa9ZCBPMocrgO63ELGCBEAo3BXPKCB0z75uvYJtqEWx_T7GXS1aB7pQDVgBMjaGmIOQ
