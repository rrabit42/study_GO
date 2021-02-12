package main

import (
	"fmt"

	"github.com/rrabit42/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
