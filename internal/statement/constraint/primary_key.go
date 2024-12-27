package constraint

import "cheetah/internal/statement/builder"

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:12:50
* @description:
**/

// PrimaryKey 主键
// 主键中所有的列必须为 NOT NULL,若没有显示声明为 NOT NULL,
// MySQL会默认声明为 NOT NULL
// 不允许自定义名字,统一为 PRIMARY
type PrimaryKey struct {
	// Columns 主键作用的列名
	Columns []string
}

func (pk PrimaryKey) BuildClause(builder builder.ClauseBuilder) {
	builder.WriteString("PRIMARY KEY (")
	for i, column := range pk.Columns {
		builder.WriteQuoteTo(column)
		if i < len(pk.Columns)-1 {
			builder.WriteByte(',')
		}
	}
	builder.WriteByte(')')
}
