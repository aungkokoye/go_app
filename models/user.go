package models

import "database/sql"

type SyncUser struct {
	UserID    int            `db:"user_id"`
	Email     string         `db:"email"`
	Username  string         `db:"username"`
	Firstname sql.NullString `db:"firstname"`
	Lastname  sql.NullString `db:"lastname"`
	Created   sql.NullString `db:"created"`
	LastLogin sql.NullString `db:"last_login"`
}
