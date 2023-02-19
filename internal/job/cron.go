package job

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/E7ast1c/Cupbearer/internal/dal"
	"github.com/jackc/pgtype"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func (h Handler) RunCron(ctx context.Context) {
	cj := cron.New(cron.WithChain(
		cron.Recover(cron.DefaultLogger), // or use cron.DefaultLogger
	), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))

	_, err := cj.AddFunc("@every 1m", h.ReadBirthdays)
	if err != nil {
		logrus.Error(err)
	}

	cj.AddFunc("@every 1m", func() { fmt.Println("Every hour thirty") })
	cj.Start()

	fmt.Println(cj.Entries())

	ch := make(chan int)
	<-ch
}

func (h Handler) ReadBirthdays() {
	db := dal.NewDBExec(h.Conn)
	db.CreateBirthday(h.Ctx, dal.Birthday{
		UserID:     123,
		PersonName: "Artem444ik",
		RemindAt: pgtype.Date{
			Time: time.Now(),
		},
		Payload: "Pozdravit",
		BirthdayDate: pgtype.Date{
			Time: time.Now(),
		},
	})
}
