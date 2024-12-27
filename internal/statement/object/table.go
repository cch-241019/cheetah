package object

import (
	"cheetah/internal/statement/builder"
	"cheetah/internal/statement/constraint"
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
}

func (tbl Table) BuildStmt(builder builder.StmtBuilder) {
	builder.WriteString("CREATE TABLE ")
	builder.WriteQuoteTo(tbl.Name + " ")
	builder.WriteByte('(')
}
