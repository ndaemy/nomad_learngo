package main

import "fmt"

func main() {
	names := []string {"ndaemy", "soultree", "fly"}
	names = append(names, "test")
	fmt.Println(names)
}
