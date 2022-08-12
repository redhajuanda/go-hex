package migration

import (
	"fmt"
	"go-hex/app"
	"go-hex/configs"
	"go-hex/pkg/db"
	"go-hex/pkg/logger"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

const (
	MIGRATION_TYPE_UP    = "up"
	MIGRATION_TYPE_DOWN  = "down"
	MIGRATION_TYPE_FRESH = "fresh"
)

type Migration struct {
	cfg *configs.Config
	log logger.Logger
	db  *bun.DB
}

func New() *Migration {
	cfg := configs.LoadDefault()
	log := logger.New(cfg.Server.NAME, app.Version)
	logger.SetFormatter(&logrus.JSONFormatter{})
	db, err := db.NewBunMariaDBConn(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName)
	if err != nil {
		panic(err)
	}
	return &Migration{
		cfg,
		log,
		db,
	}
}

func (m *Migration) Start(migrationType string) {

	m.log.Infof("start migration %s", migrationType)
	if migrationType == MIGRATION_TYPE_FRESH {
		if m.cfg.Server.ENV.IsProd() {
			m.log.Fatalf("cannot migrate fresh in production")
			return
		}
		m.log.Info("drop database")
		_, err := m.db.Exec(fmt.Sprintf("DROP DATABASE %s; CREATE DATABASE %s;", m.cfg.Database.DBName, m.cfg.Database.DBName))
		if err != nil {
			panic(err)
		}
		m.db, err = db.NewBunMariaDBConn(m.cfg.Database.Host, m.cfg.Database.Port, m.cfg.Database.Username, m.cfg.Database.Password, m.cfg.Database.DBName)
		if err != nil {
			panic(err)
		}
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./scripts/migrations/mariadb",
	}

	var direction migrate.MigrationDirection
	switch migrationType {
	case MIGRATION_TYPE_UP:
		direction = migrate.Up
	case MIGRATION_TYPE_DOWN:
		direction = migrate.Down
	case MIGRATION_TYPE_FRESH:
		direction = migrate.Up
	}

	count, err := migrate.Exec(m.db.DB, "mysql", migrations, direction)
	if err != nil {
		panic(err)
	}
	m.log.Infof("applied %d migrations", count)
}
