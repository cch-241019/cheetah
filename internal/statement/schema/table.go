package schema

import (
	"cheetah/internal/constraint"
	"cheetah/internal/object"
)

type Table struct {
	Name       string
	Engine     string
	Comment    string
	Collation  string
	Columns    []Column
	PrimaryKey constraint.PrimaryKey
	Indexes    []object.Index
}
