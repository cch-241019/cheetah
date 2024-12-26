package ddl

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:42:23
* @description:
**/

type DropType int

const (
	DropPrimaryKey DropType = iota
	DropForeignKey
	DropIndex
	DropView
	DropTable
	DropDatabase
)
