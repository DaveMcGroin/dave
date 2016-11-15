package reminder

import (
	"fmt"
	"database/sql"
	"time"
	"log"
	"github.com/zuzuleinen/dave/email"
)

func Remind(db *sql.DB) {
	reminders := Read(db)
	for _, r := range reminders {
		if shouldRemind(r) {
			sendReminderMail(r.Name)
		}
	}
}

func shouldRemind(r Reminder) bool {
	t, err := time.Parse("Monday, 2 Jan 2006 at 15:04", r.Time)

	if err != nil {
		log.Fatalln("Cannot parse time", err)
	}
	thisYear, thisMonth, today := time.Now().Date()
	y, m, d := t.Date()
	if (thisYear == y && thisMonth == m && today == d) {
		return true
	}
	return false
}

func sendReminderMail(todo string) {
	subject := fmt.Sprintf("Reminder: %s", todo)
	body := fmt.Sprintf("Hey, you need to %s", todo)

	email.Send("andrey.boar@gmail.com", subject, body)
}