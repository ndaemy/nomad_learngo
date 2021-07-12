package main

import (
	"fmt"
	"time"
)

func isCool(person string, c chan string) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- person + " is cool"
}

func main() {
	c := make(chan string)
	people := [4]string{"ndaemy", "soultree", "fly", "captain"}
	for _, person := range people {
		go isCool(person, c)
	}
	for range people {
		fmt.Println(<-c)
	}
}
