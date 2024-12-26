package ddl

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:42:53
* @description:
**/

type AddType int

const (
	AddColumn AddType = iota
	AddPrimaryKey
	AddForeignKey
	AddUnique
	AddNotNull
	AddDefault
	AddFirst
	AddAfter
	AddUniqueIndex
	AddIndex
)
