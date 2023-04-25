package remindme_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/remindme"
)

func TestRemindme(t *testing.T){
	t.Parallel()
	tmp := t.TempDir() + "/remindme2.gobdb"
	r, err := remindme.New(remindme.WithPath(tmp))
	if err != nil {
		t.Fatalf("unable to create are reminder db for path %s: %s", tmp, err)
	}
	err = r.Add("buy milk today")
	if err != nil {
		t.Fatalf("unable to add reminder: %s", err)
	}

	err = r.Add("buy cake today")
	if err != nil {
		t.Fatalf("unable to add reminder: %s", err)
	}

	err = r.Add("buy gift today")
	if err != nil {
		t.Fatalf("unable to add reminder: %s", err)
	}

	want := remindme.Reminders{
		{"buy milk today"},
		{What: "buy cake today"},
		{What: "buy gift today"},
	}

	got, err := r.List()
	if err != nil {
		t.Fatalf("unable to list reminders: %s", err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))	
	}
}

func TestRemindme2(t *testing.T){
	t.Parallel()
	tmp := t.TempDir() + "/remindme2.gobdb"
	r, err := remindme.New(remindme.WithPath(tmp))
	if err != nil {
		t.Fatalf("unable to create are reminder db for path %s: %s", tmp, err)
	}
	err = r.Add("buy milk today")
	if err != nil {
		t.Fatalf("unable to add reminder: %s", err)
	}

	err = r.Add("buy cake today")
	if err != nil {
		t.Fatalf("unable to add reminder: %s", err)
	}

	err = r.Add("buy gift today")
	if err != nil {
		t.Fatalf("unable to add reminder: %s", err)
	}

	want := remindme.Reminders{
		{"buy milk today"},
		{What: "buy cake today"},
		{What: "buy gift today"},
	}

	got, err := r.List()
	if err != nil {
		t.Fatalf("unable to list reminders: %s", err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))	
	}
}
