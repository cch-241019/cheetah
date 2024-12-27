package constraint

import (
	"strings"
	"testing"
)

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

func TestPrimaryKey(t *testing.T) {
	const pkCluase = "PRIMARY KEY (`id`,`name`,`age`)"
	pk := PrimaryKey{Columns: []string{"id", "name", "age"}}
	builder := testBuilder{&strings.Builder{}}
	pk.BuildClause(builder)
	if builder.String() != pkCluase {
		t.Error(builder.String())
	}
}
