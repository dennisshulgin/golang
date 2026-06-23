package main

import (
	"fmt"
)

type user struct {
	id    int
	name  string
	age   int
	email string
}

type product struct {
	id    int
	name  string
	price int
}

type order struct {
	id       int
	price    int
	products []product
}

func (o *order) Total() int {
	total := 0
	for _, p := range o.products {
		total += p.price
	}
	return total
}

func (u *user) Validate() bool {
	if u.id <= 0 || u.age <= 0 || u.name == "" {
		return false
	}
	return true
}

func main() {
	u := user{id: 1, name: "Denis", age: 28, email: "den@yandex.ru"}
	fmt.Println(u.Validate())
	p1 := product{id: 1, name: "Product 1", price: 100}
	p2 := product{id: 2, name: "Product 2", price: 200}
	o := order{id: 1, price: 300, products: []product{p1, p2}}
	fmt.Println(o.Total())
}
