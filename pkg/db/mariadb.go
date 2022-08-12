package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMariaDBConn(host string, port int, user, password, dbName string) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)
	db, err := sqlx.Connect("mysql", dataSource)

	if err != nil {
		return nil, err
	}

	return db, nil
}
