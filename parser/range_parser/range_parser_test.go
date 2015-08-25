package range_parser

import (
	"testing"
)

func TestNormalInput(t *testing.T) {
	opt := "order_by=a b c d;order=desc;max=100;offset=200;want=f g"
	rangOpt, err := Parse(opt)

	if err != nil {
		t.Error(err)
	}

	if rangOpt.Max != 100 {
		t.Error("max value not equal 100")
	}

	if rangOpt.Offset != 200 {
		t.Error("offset value not equal 200")
	}

	if rangOpt.OrderBy == nil {
		t.Error("range opt order key is nil")
	}

	orderKey := []string{"a", "b", "c", "d"}

	if len(orderKey) != len(rangOpt.OrderBy) {
		t.Error("range opt order key length error")
	}

	for id, item := range orderKey {
		if rangOpt.OrderBy[id] != item {
			t.Error("range opt order key error")
		}
	}

	want := []string{"f", "g"}

	if rangOpt.Want == nil || len(want) != len(rangOpt.Want) {
		t.Error("range opt want key length error")
	} else {
		for id, item := range want {
			if rangOpt.Want[id] != item {
				t.Error("range opt want key error")
			}
		}
	}

}
