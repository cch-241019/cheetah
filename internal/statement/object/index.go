package object

import (
	"cheetah/internal/builder"
	"cheetah/internal/statement"
)

type IndexType int

const (
	BTree IndexType = iota
	Hash
	Fulltext
	Spatial
	GIN
)

type Index struct {
	Name    string
	Type    IndexType
	Columns []statement.Column
}

func (idx Index) Build(builder builder.Builder) {
	switch idx.Type {
	case BTree:
		idx.buildBTree(builder)
	case Hash:
		idx.buildHash(builder)
	case Fulltext:
		idx.buildFulltext(builder)
	case Spatial:
		idx.buildSpatial(builder)
	case GIN:
		idx.buildGIN(builder)
	}
}

func (idx Index) buildBTree(builder builder.Builder) {
	builder.WriteString("INDEX ")
	if idx.Name != "" {
		builder.WriteQuoteTo(idx.Name)
	}
	builder.WriteString(" (")
	for i, column := range idx.Columns {
		builder.WriteQuoteTo(column.Name)
		if i != len(idx.Columns)-1 {
			builder.WriteByte(',')
		}
	}
	builder.WriteByte(')')
}

func (idx Index) buildFulltext(builder builder.Builder) {
	builder.WriteString("FULLTEXT INDEX ")
	if idx.Name != "" {
		builder.WriteQuoteTo(idx.Name)
	}
	builder.WriteString(" (")
	for i, column := range idx.Columns {
		builder.WriteQuoteTo(column.Name)
		if i != len(idx.Columns)-1 {
			builder.WriteByte(',')
		}
	}
	builder.WriteByte(')')
}

func (idx Index) buildHash(builder builder.Builder) {
	builder.WriteString("INDEX ")
	if idx.Name != "" {
		builder.WriteQuoteTo(idx.Name)
	}
	builder.WriteString(" USING HASH (")
	for i, column := range idx.Columns {
		builder.WriteQuoteTo(column.Name)
		if i != len(idx.Columns)-1 {
			builder.WriteByte(',')
		}
	}
}

func (idx Index) buildSpatial(builder builder.Builder) {
	builder.WriteString("SPATIAL INDEX ")
	if idx.Name != "" {
		builder.WriteQuoteTo(idx.Name)
	}
	builder.WriteString(" (")
	for i, column := range idx.Columns {
		builder.WriteQuoteTo(column.Name)
		if i != len(idx.Columns)-1 {
			builder.WriteByte(',')
		}
	}
	builder.WriteByte(')')
}

func (idx Index) buildGIN(builder builder.Builder) {

}
