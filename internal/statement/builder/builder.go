package builder

type Builder interface {
	WriteString(string)
	WriteRune(rune)
	WriteByte(byte)
	WriteQuoteTo(string)
	String() string
}

type StmtBuilder interface {
	Builder
}

type ClauseBuilder interface {
	Builder
}

type Clause interface {
	Build(Builder) error
}
