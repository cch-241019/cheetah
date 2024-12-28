package change_log

type Stmt interface {
	Rollback() error
}
