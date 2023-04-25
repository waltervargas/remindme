package remindme

import (
	"fmt"
	"os"

	"github.com/waltervargas/gobdb"
)

type Remindme struct {
	path string
	db gobdb.Gobdb
}

type Reminders []Reminder
type Reminder struct {
	What string
}

func WithPath(p string) func(r *Remindme) {
	return func(r *Remindme) { 
		r.path = p
	}
}

type opts func(r *Remindme)

func New(opts ...opts) (*Remindme, error) {
	r := &Remindme{
		// set default path
		path: os.Getenv("HOME") + "/.reminders.gobdb",
	}

	for _, o := range opts {
		o(r)
	}

	db, err := gobdb.Open(r.path, gobdb.WithType(Reminders{}))
	if err != nil {
		return nil, err
	}
	r.db = db

	return r, nil
}

func Add(s string) error {
	r, err := New()
	if err != nil {
		return err
	}
	return r.Add(s)
} 

func (r *Remindme) Add(s string) error {
	var reminders Reminders

	data := r.db.List()
	remindersData, ok := data["reminders"]
	if ok {
		reminders, ok = remindersData.(Reminders)
		if !ok {
			return fmt.Errorf("invalid data type for reminders")
		}
	}

	reminders = append(reminders, Reminder{s})
	err := r.db.Add(gobdb.Data{
		"reminders": reminders,
	})
	if err != nil {
		return err
	}
	return nil
}

func List() (Reminders, error) {
	r, err := New()
	if err != nil {
		return nil, err
	}
	return r.List()
} 

func (r *Remindme) List() (Reminders, error) {

	data := r.db.List()
	remindersData, ok := data["reminders"]
	if !ok{
		return nil, fmt.Errorf("reminders not found")
	}

	reminders, ok := remindersData.(Reminders)
	if !ok {
		return nil, fmt.Errorf("invalid data type for reminders")
	}

	return reminders, nil
}
