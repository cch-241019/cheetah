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

// ColumnMeta 列元数据
type ColumnMeta struct {
	// DataType 数据类型
	DataType string
	// CollationName 对于字符串列，为排序规则
	CollationName string
	// ColumnKey empty, PRI, UNI, MUL
	ColumnKey string
}

func ColumnMeta2Column(meta *ColumnMeta, column *Column) error {
	return nil
}
