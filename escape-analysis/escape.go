package escapeanalysis

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

/*
a, b no escape
*/
func Sum(a, b int) int {
	return a + b
}

/*
x escape
*/
func NewInt() *int {
	x := 42
	return &x
}

/*
v no escape
*/
func PrintValue(v any) {
	fmt.Println(v)
}

/*
[]int escape
*/
func MakeSmallSlice() []int {
	return make([]int, 10)
}

/*
u.Age no escape
*/
func (u User) IsAdult() bool {
	return u.Age >= 18
}

/*
u.Age no escape, but object u may be store in heap
*/
func (u *User) Birthday() {
	u.Age++
}

func NewUserValue() User {
	return User{Name: "Denis", Age: 28}
}

func NewUserPointer() *User {
	return &User{Name: "Denis", Age: 28}
}
