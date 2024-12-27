package object

import "cheetah/internal/statement/builder"

// https://dev.mysql.com/doc/refman/8.4/en/information-schema-statistics-table.html

type IndexType int

const (
	BTreeIndex IndexType = iota
	FulltextIndex
	HashIndex
	RTreeIndex
)

type AlgorithmType int

const (
	DefalutAlgorithm AlgorithmType = iota
	InplaceAlgorithm
	CopyAlgorithm
)

type LockType int

const (
	DefaultLock LockType = iota
	NoneLock
	SharedLock
	ExclusiveLock
)

type IndexColumn struct {
	Name string
	// SubPart 索引前缀
	// 整个列都被索引为"",部分被索引则为
	// 索引字符数
	SubPart int
	// Expression 表达式
	Expression string
}

type Index struct {
	// Table 表名称
	Table string
	// Name 索引名称，若索引为主键
	// 元数据表中为PRIMARY
	Name string
	// Unique 索引能否包含重复项
	Unique bool
	// Columns 索引作用的列
	Columns []IndexColumn
	// Type 索引类型
	Type IndexType
	// 索引注释
	Comment string
	// Visible 索引是否对优化器可见
	Visible bool
	// KeyBlockSize 指定索引页的大小
	KeyBlockSize  int
	LockType      LockType
	AlgorithmType AlgorithmType
	// WithParser 指定 Fulltext使用的解析器
	WithParser string
	// EngineAttribute
	EngineAttribute string
	// SecondaryEngineAttribute
	SecondaryEngineAttribute string
}

// https://dev.mysql.com/doc/refman/8.4/en/create-index.html

func (idx Index) BuildStmt(builder builder.StmtBuilder) {

}

func (idx Index) BuildClause(builder builder.ClauseBuilder) {

}
