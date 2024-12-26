package object

import "cheetah/internal/schema"

type Index struct {
	Name   string
	Column []schema.Column
}
