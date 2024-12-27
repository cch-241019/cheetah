package object

import (
	"cheetah/internal/statement/constraint"
)

type Table struct {
	Name       string
	Engine     string
	Comment    string
	Collation  string
	Columns    []Column
	PrimaryKey constraint.PrimaryKey
	Indexes    []Index
}
