package main

import (
	"fmt"

	"example.com/accounts"
)

func main() {
	account := accounts.NewAccount("ndaemy")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
