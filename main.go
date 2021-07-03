package main

import "fmt"

type person struct {
	name      string
	age       int
	favSinger []string
}

func main() {
	singers := []string{"PHS", "IU"}
	ndaemy := person{name: "ndaemy", age: 25, favSinger: singers}
	fmt.Println(ndaemy)
}
