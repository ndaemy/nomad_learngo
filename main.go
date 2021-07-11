package main

import (
	"fmt"
	"time"
)

func isCool(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}

func main() {
	c := make(chan bool)
	people := [2]string{"ndaemy", "soultree"}
	for _, person := range people {
		go isCool(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}
