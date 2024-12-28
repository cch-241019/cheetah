package types

import (
	"cheetah/internal/statement/builder"
	"database/sql"
	"strconv"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:26:37
* @description:
**/

type Int struct {
	Numeric
	Fullname string
	Default  sql.NullInt32
}

func (i *Int) Name() string {
	return "INT"
}

func (i *Int) Build(builder builder.Builder) error {
	builder.WriteString(i.Fullname)
	if i.Unsigned {
		builder.WriteString(" UNSIGNED")
	}
	if i.Default.Valid {
		builder.WriteString(" DEFAULT ")
		builder.WriteByte('\'')
		builder.WriteString(strconv.Itoa(int(i.Default.Int32)))
		builder.WriteByte('\'')
	}
	return nil
}

func (i *Int) WriteFullname(fullname string) {
	i.Fullname = fullname
}

func (i *Int) Reset() {
	*i = Int{}
}
