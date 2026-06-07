package main

import "fmt"



func main() {
	var s []string = make([]string, 3)
	s[0] = "1"
	s[1] = "2"
	s[2] = "3"
	var s2 []string = s[0:1]
	
	fmt.Println(s)
	fmt.Println(s2)
	s2[0] = "4"
	fmt.Println(s)
	fmt.Println(s2)

}
