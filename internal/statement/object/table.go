package object

import (
	"cheetah/internal/statement/builder"
	"cheetah/internal/statement/constraint"
	"database/sql"
)

type TableType int

const (
	Default TableType = iota // TABLE
	View
	SystemView
)

// https://dev.mysql.com/doc/refman/9.1/en/information-schema-tables-table.html

type Table struct {
	// Name 表名
	Name string
	// AutoIncrement 自增值
	AutoIncrement sql.NullInt32
	// Columns 列
	Columns []Column
	// PrimaryKey 主键
	// 主键中所有的列必须为 NOT NULL,若没有显示声明为 NOT NULL,
	// MySQL会默认声明为 NOT NULL
	// 不允许自定义名字,统一为 PRIMARY
	PrimaryKey constraint.PrimaryKey
	// Indexes 索引
	Indexes []Index
	// Engine 表使用给的引擎
	Engine string
	// Comment 注释
	Comment string
	// Collation 编码
	Collation string
}

func (tbl *Table) Build(builder builder.Builder) error {
	builder.WriteString("CREATE TABLE ")
	builder.WriteQuoteTo(tbl.Name)
	builder.WriteByte('(')
	builder.WriteByte('\n')

	if tbl.Columns != nil && len(tbl.Columns) > 0 {
		for i, column := range tbl.Columns {
			if column.PrimaryKey {
				tbl.PrimaryKey.Columns = append(tbl.PrimaryKey.Columns, column.Name)
			}
			err := column.Build(builder)
			if err != nil {
				return err
			}
			if i < len(tbl.Columns)-1 {
				builder.WriteString(",\n")
			}
		}
	}
	if len(tbl.PrimaryKey.Columns) > 0 {
		builder.WriteString(",\n")
		err := tbl.PrimaryKey.Build(builder)
		if err != nil {
			return err
		}
	}
	if tbl.Indexes != nil {
		for _, index := range tbl.Indexes {
			index.Build(builder)
		}
	}
	builder.WriteString("\n)")
	if tbl.Engine != "" {
		builder.WriteString(" ENGINE=")
		builder.WriteString(tbl.Engine)
	}
	if tbl.Collation != "" {
		builder.WriteString(" DEFAULT CHARSET=")
		builder.WriteString(tbl.Collation)
	}
	if tbl.AutoIncrement.Int32 != 0 {
		builder.WriteString(" AUTO_INCREMENT=")
		builder.WriteString(string(tbl.AutoIncrement.Int32))
	}
	return nil
}
