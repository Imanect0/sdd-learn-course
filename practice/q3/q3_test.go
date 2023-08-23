package q3

import "testing"

func TestGetPersonInfo(t *testing.T) {
	p1 := Person{"Taro", 20}
	expected1 := "Name is Taro, Age is 20"
	if result1 := GetPersonInfo(p1); result1 != expected1 {
		t.Errorf("GetPersonInfo(%v) = %v, expected %v", p1, result1, expected1)
	}

	p2 := Person{"Jiro", 30}
	expected2 := "Name is Jiro, Age is 30"
	if result2 := GetPersonInfo(p2); result2 != expected2 {
		t.Errorf("GetPersonInfo(%v) = %v, expected %v", p2, result2, expected2)
	}
}