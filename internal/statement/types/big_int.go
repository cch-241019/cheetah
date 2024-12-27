package types

import (
	"cheetah/internal/statement/builder"
	"strconv"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:26:56
* @description:
**/

type BigInt struct {
	Numeric
}

func (b BigInt) Build(builder builder.ClauseBuilder) {
	builder.WriteString("BIGINT")
	if b.DisplayWidth > 0 {
		builder.WriteByte('(')
		builder.WriteString(strconv.Itoa(b.DisplayWidth))
		builder.WriteByte(')')
	}
	if b.Unsigned {
		builder.WriteString("UNSIGNED ")
	}
}
