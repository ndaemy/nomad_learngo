package main

import (
	"fmt"
	"time"
)

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is Good", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go count("ndaemy")
	go count("soultree")
	time.Sleep(time.Second * 5)
}
