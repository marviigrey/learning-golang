package main

import "testing"

func TestCheckProcess(t *testing.T) { //always start your test function with a capital letter.
	i := 13
	expected := "YES"
	res := checkProcess(i)
	if expected != res {
		t.Errorf("expected the value %v and got: %v", expected, res)
	}
}
