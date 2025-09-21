package db

import (
	"fmt"

	"github.com/gowaliullah/ecommerce/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf config.DBConfig) string {

	connString := fmt.Sprintf(
		"user=%c password=%c host=%c port=%c dbname=%c",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)

	if !cnf.EnableSSLMODE {
		connString += " sslmode=disable"
	}
	return connString
}

func NewConnection(cnf config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
