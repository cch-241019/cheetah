package constraints

/*
* @author: Chen Chiheng
* @date: 2024/12/25 21:12:50
* @description:
**/

// PRIMARY KEY

type PrimaryKey struct {
	// Name 自定义的主键名称
	Name string

	// Columns 主键作用的列名
	Columns []string
}

// 删除主键 DROP PRIMARY KEY

func DropPrimaryKey() {

}
