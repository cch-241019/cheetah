package statement

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:11:07
* @description:
**/

// Column is a database schema column。
type Column struct {
	Name            string
	OrdinalPosition int
	AutoIncrement   bool
	Nullable        bool
	Default         string
	Type            string
	Comment         string
}

type ColumnMeta struct {
	Catalog   string
	Schema    string
	TableName string
	Name      string
	Position  int
	Default   string
	Nullable  bool
	Type      string
	// MaximumLength 最大长度
	MaximumLength int
	// OctetLength 最大字节长度
	OctetLength       int
	NumericPrecision  int
	NumericScale      int
	DateTimePrecision int
	CharacterSet      string
	Collation         string
	ColumnType        string
	ColumnKey         string
	Extra             string
	Privileges        string
	Comment           string
	GenerationExpr    string
	SrsID             int
}
