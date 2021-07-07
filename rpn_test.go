package main

import "testing"

/*
func TestEx01(t *testing.T) {
	expect := 3
	slice := []string{"1", "2", "+"}
	result := calculater(slice)
	if result != expect {
		t.Error("Test01 is failed")
	}

}

func TestEx02(t *testing.T) {
	expect := 1
	slice := []string{"2", "1", "-"}
	result := calculater(slice)
	if result != expect {
		t.Error("Test02 is failed")
	}
}

func TestEx03(t *testing.T) {
	expect := 6
	slice := []string{"3", "2", "*"}
	result := calculater(slice)
	if result != expect {
		t.Error("Test03 is failed")
	}
}
*/
func TestEx04(t *testing.T) {
	expect := 2
	slice := []string{"3", "2", "*"}
	result := search(slice)
	if result != expect {
		t.Error("Test04 is failed")
	}
}
