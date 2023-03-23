package remindme

import (
	"fmt"
	"os"

	"github.com/waltervargas/gobdb"
)

type Reminder struct {
	What string
	When string
}

func Add(s string) error {
	db, err := gobdb.Open(os.Getenv("HOME") + "/.reminders.gobdb")
	if err != nil {
		return err
	}
	data := gobdb.Data{
		"reminders": []Reminder{{What: s}},
	}
	err = db.Add(data)
	if err != nil {
		return err
	}
	return nil
}

func List() ([]Reminder, error) {
	db, err := gobdb.Open(os.Getenv("HOME") + "/.reminders.gobdb")
	if err != nil {
		return nil, err
	}

	data := db.List()
	remindersData, ok := data["reminders"]
	if !ok{
		return nil, fmt.Errorf("reminders not found")
	}

	reminders, ok := remindersData.([]Reminder)
	if !ok {
		return nil, fmt.Errorf("invalid data type for reminders")
	}

	return reminders, nil
}
