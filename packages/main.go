package main

import (
	"fmt"
	"example/packages/mathutils"
)

func main() {
	num := mathutils.Add(1, 2)
	fmt.Println(num)
}