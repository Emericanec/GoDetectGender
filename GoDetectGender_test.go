package GoDetectGender

import (
	"testing"
	"fmt"
)

func TestGetGender(t *testing.T) {
	maleFullName := FullName{"Светличный", "Андрей", "Андреевич"}
	maleResult := GetGender(maleFullName)
	fmt.Println(maleResult)
	assertEqual(t, maleResult, MALE, "")

	femaleFullName := FullName{"Светличная", "Татьяна", "Сергеевна"}
	femaleResult := GetGender(femaleFullName)
	fmt.Println(femaleResult)
	assertEqual(t, femaleResult, FEMALE, "")
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
