package types

import (
	"strings"
	"testing"
)

const testIntClause1 = "INT(10) UNSIGNED"

func TestInt(t *testing.T) {
	typ := Int{
		Fullname: "INT(10)",
		Numeric:  Numeric{Unsigned: true},
	}
	builder := testBuilder{
		&strings.Builder{},
	}
	err := typ.Build(builder)
	if err != nil {
		t.Fatal(err)
	}
	if builder.String() != testIntClause1 {
		t.Fatal(builder.String())
	}
}
