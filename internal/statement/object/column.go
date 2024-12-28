package object

import (
	"cheetah/internal/statement/builder"
	"cheetah/internal/statement/types"
	"strings"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:11:07
* @description:
**/

// Column is a database schema column。
type Column struct {
	// Name 列名
	Name string
	// DataType 数据类型
	// 默认值绑定在数据类型上
	DataType string
	// ColumnType 完整数据类型
	ColumnType string
	// Nullable 能否为NUll
	Nullable bool
	// Visible 可见性
	Visible bool
	// AutoIncrement 自增
	AutoIncrement bool
	// PrimaryKey 是否为主键
	// constraint.PrimaryKey 对主键进行更为详细的描述
	PrimaryKey bool
	// Comment 注释
	Comment string
	// MySQL 8.4 已忽略
	// ColumnFormat 格式
	// FIXED DYNAMIC DEFAULT
	ColumnFormat string
	// EngineAttribute 引擎属性
	EngineAttribute string
	// Storage 存储类型
	// 非NDB表不生效
	Storage string
	// OrdinalPosition 列顺序
	OrdinalPosition int
}

func (col Column) Build(builder builder.Builder) error {
	builder.WriteString(col.Name)
	builder.WriteByte(' ')
	dataType, err := types.GetDataTypeByName(strings.ToUpper(col.DataType))
	if err != nil {
		return err
	}
	dataType.WriteFullname(strings.ToUpper(col.ColumnType))
	err = dataType.Build(builder)
	if err != nil {
		return err
	}
	return nil
}
