package main

import "fmt"

func main() {
	ndaemy := map[string]string{"name": "ndaemy", "age": "25"}
	for _, value := range ndaemy {
		fmt.Println(value)
	}
}
