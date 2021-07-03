package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 10
	fmt.Println(a, *b)
}
