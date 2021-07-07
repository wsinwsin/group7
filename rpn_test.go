package main

import "testing"

func TestEx01(t *testing.T) {
	expect := "3"
	result := calculater("1", "2", "+")
	if result != expect {
		t.Error("Test01 is failed")
	}

}

func TestEx02(t *testing.T) {
	expect := "1"
	result := calculater("2", "1", "-")
	if result != expect {
		t.Error("Test02 is failed")
	}
}

func TestEx03(t *testing.T) {
	expect := "6"
	result := calculater("3", "2", "*")
	if result != expect {
		t.Error("Test03 is failed")
	}
}

func TestEx04(t *testing.T) {
	expect := 3
	slice := []string{"3", "2", "5", "*"}
	result := search(slice)
	if result != expect {
		t.Error("Test04 is failed")
	}
}
