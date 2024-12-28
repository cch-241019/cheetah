package types

import (
	"cheetah/internal/statement/builder"
	"database/sql"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:29:36
* @description:
**/

type Varchar struct {
	// Characterset 字符编码
	Characterset string
	// Collate 排序顺序
	Collate string
	// Default 默认值
	Default sql.NullString
	// Fullname 类型全名
	Fullname string
}

func (v *Varchar) Name() string {
	return "VARCHAR"
}

func (v *Varchar) Build(builder builder.Builder) error {
	builder.WriteString(v.Fullname)
	if v.Characterset != "" {
		builder.WriteString(" CHARACTER SET ")
		builder.WriteString(v.Characterset)
	}
	if v.Collate != "" {
		builder.WriteString(" COLLATE ")
		builder.WriteString(v.Collate)
	}
	if v.Default.Valid {
		builder.WriteString(" DEFAULT ")
		builder.WriteString(v.Default.String)
	}
	return nil
}

func (v *Varchar) WriteFullname(fullname string) {
	v.Fullname = fullname
}

func (v *Varchar) Reset() {
	*v = Varchar{}
}
