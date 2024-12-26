package schema

import (
	"cheetah/internal/statement"
	"cheetah/internal/statement/constraint"
	"cheetah/internal/statement/object"
)

type Table struct {
	Name       string
	Engine     string
	Comment    string
	Collation  string
	Columns    []statement.Column
	PrimaryKey constraint.PrimaryKey
	Indexes    []object.Index
}
