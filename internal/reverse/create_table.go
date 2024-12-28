package reverse

import (
	"cheetah/internal/statement/object"
	"database/sql"
)

type RevType int

const (
	tableBaseStmt = `select 
    engine, 
    auto_increment, 
    table_comment, 
    table_collation 
	from information_schema.tables 
	where table_schema = ? and table_name = ?`
)

func RevCreateTableStmt(db *sql.DB, schemaName, tableName string) (*object.Table, error) {
	row := db.QueryRow(tableBaseStmt, schemaName, tableName)
	table := object.Table{}
	err := row.Scan(
		&table.Engine,
		&table.AutoIncrement,
		&table.Collation,
		&table.Comment,
	)
	if err != nil {
		return nil, err
	}
	table.Name = tableName
	return &table, nil
}

const queryColumnMetaStmt = `select
column_name,
ordinal_position,
is_nullable,
column_default,
data_type,
character_maximum_length,
numeric_precision,
numeric_scale,
datetime_precision,
column_type,
column_comment,
column_key,
extra,
from information_schema.columns 
where table_schema = ? and table_name = ?`

func queryColumn(db *sql.DB, schemaName, tableName string) {

}
