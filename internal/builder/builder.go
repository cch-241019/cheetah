package builder

type Builder interface {
	WriteString(string)
	WriteRune(rune)
	WriteByte(byte)
	WriteQuoteTo(string)
	String() string
}
