package remindme_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/remindme"
)

func TestRemindme(t *testing.T){
	err := remindme.Add("buy milk today")
	if err != nil {
		t.Fatalf("unable to add reminder: %s", err)
	}

	want := []remindme.Reminder{
		{What: "buy milk today"},
	}

	got, err := remindme.List()
	if err != nil {
		t.Fatalf("unable to list reminders: %s", err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))	
	}
}
