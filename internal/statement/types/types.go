package types

import "cheetah/internal/statement/builder"

/*
* @author: Chen Chiheng
* @date: 2024/12/25 22:09:40
* @description:
**/

type DataType interface {
	builder.Clause
	Default(builder.ClauseBuilder)
}

// Numeric 数值类型
type Numeric struct {
	// Unsigned 是否是无符号的
	Unsigned bool
	// DisplayWidth 显示宽度
	DisplayWidth int
	// 不推荐使用
	// Zerofill 若此值为true则不足显示宽度将填充零
	// 若为数值列指定ZEROFILL，MySQL会自动添加UNSIGNED 属性
	Zerofill bool
}
