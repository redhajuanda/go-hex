package mysql

import "github.com/uptrace/bun"

// DBI is a DB interface implemented by *DB and *Tx.
type DBI interface {
	NewValues(model interface{}) *bun.ValuesQuery
	NewSelect() *bun.SelectQuery
	NewInsert() *bun.InsertQuery
	NewUpdate() *bun.UpdateQuery
	NewDelete() *bun.DeleteQuery
	NewCreateTable() *bun.CreateTableQuery
	NewDropTable() *bun.DropTableQuery
	NewCreateIndex() *bun.CreateIndexQuery
	NewDropIndex() *bun.DropIndexQuery
	NewTruncateTable() *bun.TruncateTableQuery
	NewAddColumn() *bun.AddColumnQuery
	NewDropColumn() *bun.DropColumnQuery
}
