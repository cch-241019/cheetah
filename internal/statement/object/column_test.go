package object

import (
	"strings"
	"testing"
)

func TestColumn(t *testing.T) {
	t.Run("testColumnBuild", testColumnBuild)
}

func testColumnBuild(t *testing.T) {
	builder := testBuilder{&strings.Builder{}}
	col := Column{
		Name:         "id",
		DataType:     "varchar",
		ColumnType:   "varchar(20)",
		Nullable:     false,
		Visible:      true,
		Comment:      "user name",
		Characterset: "utf8mb4",
		Collate:      "utf8mb4_bin",
	}
	if err := col.Build(builder); err != nil {
		t.Fatal(err)
	}
	t.Log(strings.ToLower(builder.String()))
}
