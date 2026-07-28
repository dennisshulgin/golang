package pointers

type User struct {
	Name string
	Age  int
}

func RenameByValue(u User, name string) {
	u.Name = name
}

func RenameByPointer(u *User, name string) {
	u.Name = name
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u *User) Birthday() {
	u.Age++
}

func (u *User) SafeBirthday() bool {
	if u == nil {
		return false
	} else {
		u.Birthday()
		return true
	}
}
