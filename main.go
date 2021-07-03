package main

import (
	"fmt"

	"example.com/accounts"
)

func main() {
	account := accounts.NewAccount("ndaemy")
	// account.balance = 100000
	fmt.Println(*account)
}
