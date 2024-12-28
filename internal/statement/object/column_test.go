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
	col := Column{Name: "id", DataType: "int", ColumnType: "int(20)"}
	if err := col.Build(builder); err != nil {
		t.Fatal(err)
	}
	t.Log(builder.String())
}
