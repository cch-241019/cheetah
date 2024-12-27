package object

import (
	"cheetah/internal/statement/builder"
	"cheetah/internal/statement/types"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:11:07
* @description:
**/

// Column is a database schema column。
type Column struct {
	Name            string
	PrimaryKey      bool
	AutoIncrement   bool
	OrdinalPosition int
	Nullable        bool
	Type            types.DataType
	Comment         string
}

func (col Column) Build(builder builder.ClauseBuilder) {
	builder.WriteQuoteTo(col.Name + " ")
	col.Type.Build(builder)
	if !col.Nullable {
		builder.WriteString(" NOT NULL")
	}
	if col.AutoIncrement {
		builder.WriteString(" AUTO_INCREMENT")
	}
	if !col.Nullable {
		col.Type.Default(builder)
	}
	if col.Comment != "" {
		builder.WriteString(" COMMENT ")
		builder.WriteByte('\'')
		builder.WriteString(col.Comment)
		builder.WriteByte('\'')
	}
}

type ColumnMeta struct {
	Catalog   string
	Schema    string
	TableName string
	Name      string
	Position  int
	Default   string
	Nullable  bool
	Type      string
	// MaximumLength 最大长度
	MaximumLength int
	// OctetLength 最大字节长度
	OctetLength       int
	NumericPrecision  int
	NumericScale      int
	DateTimePrecision int
	CharacterSet      string
	Collation         string
	ColumnType        string
	ColumnKey         string
	Extra             string
	Privileges        string
	Comment           string
	GenerationExpr    string
	SrsID             int
}
