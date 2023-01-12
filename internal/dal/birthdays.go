package dal

import (
	"context"

	"github.com/jackc/pgtype"
)

type Birthday struct {
	ID           int64
	UserID       int32
	PersonName   string
	RemindAt     pgtype.Date
	Payload      string
	BirthdayDate pgtype.Date
}

func (c *DBExec) CreateBirthday(ctx context.Context, bd Birthday) {
	c.conn.Exec(
		ctx,
		"insert into $1 (user_id, person_name, remind_at, payload, birthday_date) "+
			"values ($2, $3, $4, $5, $6)",
		birthdayTable, bd.UserID, bd.PersonName, bd.RemindAt, bd.Payload, bd.BirthdayDate,
	)
}
