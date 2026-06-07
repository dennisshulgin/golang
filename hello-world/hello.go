package main

import "fmt"
import "rsc.io/quote"

func GetQuote() string {
	return quote.Go()
}

func main() {
	fmt.Println(GetQuote())
}
