package builder

type StmtBuilder interface {
	WriteString(string)
	WriteRune(rune)
	WriteByte(byte)
	WriteQuoteTo(string)
	String() string
}

type ClauseBuilder interface {
	StmtBuilder
}
