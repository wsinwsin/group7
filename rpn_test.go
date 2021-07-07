package main

import "testing"

func TestEx01(t *testing.T) {
	expect := 3
	slice := []string{"1", "2", "+"}
	result := add(slice)
	if result != expect {
		t.Error("Test01 is failed")
	}
}
