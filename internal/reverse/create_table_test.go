package reverse

import (
	"cheetah/internal/database"
	"database/sql"
	"testing"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/27 21:27:43
* @description:
**/

var testDB *sql.DB

func setupDB(t *testing.T) {
	var err error
	testDB, err = database.Create(&database.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "root",
	})
	if err != nil {
		t.Error("Failed to create test database")
	}
}

func createTable(t *testing.T) {
	setupDB(t)
}

const (
	testSchemaName = "cheetah"
	testTableName  = "test_create_table"
)

func TestRevCreateTable(t *testing.T) {
	setupDB(t)
	table, err := RevCreateTableStmt(testDB, testSchemaName, testTableName)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", table)
}
