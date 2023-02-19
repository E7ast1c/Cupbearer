package dal

import (
	"github.com/jackc/pgx/v4"
)

const (
	schema        = "public."
	birthdayTable = schema + "birthdays"
)

type DBExec struct {
	conn pgx.Conn
}

func NewDBExec(conn *pgx.Conn) *DBExec {
	return &DBExec{conn: *conn}
}
