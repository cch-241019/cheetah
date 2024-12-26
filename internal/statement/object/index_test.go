package object

import (
	"cheetah/internal/statement"
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

func (t testBuilder) String() string {
	return t.Builder.String()
}

func (t testBuilder) WriteQuoteTo(s string) {
	t.Builder.WriteByte('`')
	t.Builder.WriteString(s)
	t.Builder.WriteByte('`')
}

func TestIndex(t *testing.T) {
	t.Run("Index.Build", testIndexBuild)
}

func testIndexBuild(t *testing.T) {

	const indexClause = "INDEX `idx_test` (`col1`,`col2`,`col3`)"
	builder := testBuilder{&strings.Builder{}}
	idx := Index{
		Name: "idx_test",
		Type: BTree,
		Columns: []statement.Column{
			statement.Column{Name: "col1"},
			statement.Column{Name: "col2"},
			statement.Column{Name: "col3"},
		},
	}
	idx.Build(&builder)
	if builder.String() != indexClause {
		t.Error(builder.String(), indexClause)
	}
}
