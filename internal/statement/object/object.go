package object

import "strings"

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
