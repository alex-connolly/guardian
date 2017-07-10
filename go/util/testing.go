package util

import "testing"

func Assert(t *testing.T, condition bool, err string) {
	if !condition {
		t.Log(err)
		t.Fail()
	}
}

func AssertNow(t *testing.T, condition bool, err string) {
	if !condition {
		t.Log(err)
		t.FailNow()
	}
}
