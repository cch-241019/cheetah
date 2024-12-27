package change_log

import (
	"time"
)

type ChangeUnit struct {
	id            int
	author        string
	DateExecuted  time.Time
	OrderExecuted time.Time
	Md5Sum        string
	Description   string
	Comments      string
	Tag           string
}
