package ddl

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:46:09
* @description:
**/

type CreateType int

const (
	CreateTable CreateType = iota
	CreateView
	CreateProcedure
	CreateIndex
)

var CreateClauses = map[CreateType]string{
	CreateTable:     "CREATE TABLE",
	CreateView:      "CREATE VIEW",
	CreateProcedure: "CREATE PROCEDURE",
	CreateIndex:     "CREATE INDEX",
}
