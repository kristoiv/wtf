package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/bolt/internal"
)

func TestMarshalItem(t *testing.T) {
	now := time.Now().UTC()
	item := wtf.Item{
		ID:      "ID",
		Title:   "TITLE",
		Created: now,
		Changed: time.Time{},
		Checked: false,
	}

	var other wtf.Item
	if data, err := internal.MarshalItem(&item); err != nil {
		t.Fatal(err)
	} else if err := internal.UnmarshalItem(data, &other); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(item, other) {
		t.Fatalf("unexpected copy: %#v\nIS NOT EQUAL TO\n%#v", item, other)
	}
}
