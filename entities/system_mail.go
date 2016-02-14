package entities

import "time"

const
(
	SYSTEM_MAIL_VALID_YES = 1
	SYSTEM_MAIL_VALID_NO = 0
)
type SystemMail struct {
	Id        int64
	UserName  string
	From      string
	Password  string
	Host      string
	Port      int
	Valid     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

