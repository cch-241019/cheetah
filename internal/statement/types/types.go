package types

import (
	"cheetah/internal/statement/builder"
	"fmt"
	"strings"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/25 22:09:40
* @description:
**/

type DataType interface {
	builder.Clause
	Name() string
	WriteFullname(string)
	Reset()
}

var dataTypes = map[string]DataType{
	"INT":     &Int{},
	"VARCHAR": Varchar{},
}

// Numeric 数值类型
type Numeric struct {
	// Unsigned 是否是无符号的
	Unsigned bool
	// 不推荐使用
	// Zerofill 若此值为true则不足显示宽度将填充零
	// 若为数值列指定ZEROFILL，MySQL会自动添加UNSIGNED 属性
	Zerofill bool
}

func GetDataTypeByName(name string) (DataType, error) {
	typ, isExist := dataTypes[name]
	if !isExist {
		return nil, fmt.Errorf("unknown data type '%s'", name)
	}
	typ.Reset()
	return typ, nil
}

type testBuilder struct {
	*strings.Builder
}

func (t testBuilder) WriteString(s string) {
	t.Builder.WriteString(s)
}

func (t testBuilder) WriteRune(r rune) {
	t.Builder.WriteRune(r)
}

func (t testBuilder) WriteByte(b byte) {
	t.Builder.WriteByte(b)
}

func (t testBuilder) WriteQuoteTo(s string) {
	t.Builder.WriteByte('`')
	t.Builder.WriteString(s)
	t.Builder.WriteByte('`')
}
