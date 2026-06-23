package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	u1 := user{id: 1, name: "Denis", age: -1, email: "dens@yandex.ru"}
	u2 := user{id: 2, name: "Denis", age: 28, email: "dens@yandex.ru"}
	if u1.Validate() {
		t.Error("Expected false")
	}

	if !u2.Validate() {
		t.Error("Expected true")
	}
}

func TestTotal(t *testing.T) {
	p1 := product{id: 1, name: "1", price: 15}
	p2 := product{id: 2, name: "2", price: 35}
	o := order{id: 1, price: 10, products: []product{p1, p2}}
	if o.Total() != 50 {
		t.Error("Expected ", 50)
	}
}
