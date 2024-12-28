package object

import (
	"database/sql"
	"strings"
	"testing"
)

func TestTable(t *testing.T) {
	t.Run("testTableBuild", testTableBuild)
}

func testTableBuild(t *testing.T) {
	builder := testBuilder{&strings.Builder{}}
	tbl := &Table{
		Engine:    "InnoDB",
		Collation: "utf8",
		Name:      "test_create_table",
		Columns: []Column{{
			Name:          "id",
			DataType:      "int",
			ColumnType:    "int(20)",
			Nullable:      false,
			Visible:       true,
			Comment:       "user id",
			PrimaryKey:    true,
			ColumnDefault: sql.NullString{"0", true},
		},
			{
				Name:         "name",
				DataType:     "varchar",
				ColumnType:   "varchar(20)",
				Nullable:     false,
				Visible:      true,
				Comment:      "user name",
				Characterset: "utf8mb4",
				Collate:      "utf8mb4_bin",
				PrimaryKey:   false,
			},
		},
	}
	if err := tbl.Build(builder); err != nil {
		t.Fatal(err)
	}
	t.Log(strings.ToLower(builder.String()))
}
