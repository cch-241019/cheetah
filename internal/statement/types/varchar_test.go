package types

import (
	"strings"
	"testing"
)

const testVarcharClause = "VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin"

func TestVarchar(t *testing.T) {
	varchar := Varchar{
		Characterset: "utf8mb4",
		Collate:      "utf8mb4_bin",
		Fullname:     "VARCHAR(10)",
	}
	builder := testBuilder{&strings.Builder{}}
	err := varchar.Build(builder)
	if err != nil {
		t.Fatal(err)
	}
	if builder.String() != testVarcharClause {
		t.Error(builder.String())
	}
}
