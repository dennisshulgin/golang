package pointers

import (
	"testing"
)

func TestRenameByValue(t *testing.T) {
	u1 := User{Name: "Denis", Age: 18}
	RenameByValue(u1, "Ivan")

	if u1.Name != "Denis" {
		t.Errorf("Expected %v, but got %v", "Denis", u1.Name)
	}
}

func TestRenameByPointer(t *testing.T) {
	u1 := User{Name: "Denis", Age: 18}
	RenameByPointer(&u1, "Ivan")

	if u1.Name != "Ivan" {
		t.Errorf("Expected %v, but got %v", "Ivan", u1.Name)
	}
}

func TestIsAdult(t *testing.T) {
	u1 := User{Name: "Denis", Age: 18}
	u2 := User{Name: "Ivan", Age: 10}

	if !u1.IsAdult() {
		t.Errorf("Expected %v, but got %v", true, false)
	}

	if u2.IsAdult() {
		t.Errorf("Expected %v, but got %v", false, true)
	}
}

func TestBirthday(t *testing.T) {
	u1 := User{Name: "Denis", Age: 18}
	u1.Birthday()

	if u1.Age != 19 {
		t.Errorf("Expected %v, but got %v", 19, u1.Age)
	}
}

func TestSafeBirthday(t *testing.T) {
	u1 := User{Name: "Denis", Age: 18}
	result1 := u1.SafeBirthday()

	if !result1 {
		t.Errorf("Expected %v, but got %v", true, false)
	}

	if u1.Age != 19 {
		t.Errorf("Expected %v, but got %v", 19, u1.Age)
	}

	var u2 *User = nil
	result2 := u2.SafeBirthday()

	if result2 {
		t.Errorf("Expected %v, but got %v", false, result2)
	}
}
