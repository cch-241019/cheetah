package types

import (
	"cheetah/internal/statement/builder"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:26:37
* @description:
**/

type Int struct {
	Numeric
	Fullname string
}

func (i *Int) Name() string {
	return "INT"
}

func (i *Int) Build(builder builder.Builder) error {
	builder.WriteString(i.Fullname)
	if i.Unsigned {
		builder.WriteString(" UNSIGNED")
	}
	return nil
}

func (i *Int) WriteFullname(fullname string) {
	i.Fullname = fullname
}

func (i *Int) Reset() {
	*i = Int{}
}
