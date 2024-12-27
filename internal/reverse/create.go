package reverse

import (
	"cheetah/internal/statement/object"
	"database/sql"
)

type RevType int

func RevCreateTable(db *sql.DB, schema, table string) (*object.Table, error) {
	return nil, nil
}
