package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done.")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, uppercase := lenAndUpper("ndaemy")
	fmt.Println(totalLength, uppercase)
}
