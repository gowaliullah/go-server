package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string {
	return "user=postgres host=localhost port=5432 dbname=goserver"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
