package main

import (
	"fmt"

	"github.com/rrabit42/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(20)

	// 고는 이렇게 항상 에러를 체크하도록 강제시킴(exception이 없으니까)
	if err != nil {
		// Println()을 호출하고 프로그램을 종료시킴
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account) // Go가 자동으로 struct에서 호출해주는 method가 있음(python의 __str__ 처럼) -> String method가 있음!
}
