package main

import (
	"fmt"
	"log"

	"example.com/accounts"
)

func main() {
	account := accounts.NewAccount("ndaemy")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
