package db

import (
	"database/sql"
	"fmt"
	"go-hex/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewBunMySQLConn(env configs.Env, host string, port string, user, password, dbName string) (*bun.DB, error) {

	connection := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&multiStatements=true", user, password, host, port, dbName)
	sqldb, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open connection")
	}

	// Create a Bun db on top of it.
	db := bun.NewDB(sqldb, mysqldialect.New())

	if env.IsLocal() {
		// Print all queries to stdout.
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return db, nil
}
