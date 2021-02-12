package main

import (
	"fmt"

	"github.com/rrabit42/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)

	// 고는 이렇게 항상 에러를 체크하도록 강제시킴(exception이 없으니까)
	if err != nil {
		// Println()을 호출하고 프로그램을 종료시킴
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
