package accounts

import (
	"errors"
	"fmt"
)

// 컴파일을 하지 않을 거라 main function을 가지지 않음
// main을 컴파일 할거고, main은 banking을 사용할 것

// 소문자여서 외부에서 접근 불가, export하고 싶으면 대문자로 바꿔주기
// struct 뿐만 아니라 안의 원소 변수들도!
// ** private, public이 그냥 대소문자로 구분됨

// account struct
type Account struct {
	owner   string
	balance int
}

var errnoMoney = errors.New("Can't withdraw")

// Go의 패턴
// constructor가 없기 때문에 function으로 construct하거나 struct를 만들도록 한다.
// obejct을 만들어서 return 시키는 함수

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // object 자체를 return, 복사하는걸 원하지 않아서!
}

// method: Account struct는 method를 가지게 됨
// 보통 method의 receiver를 작성할 때, struct의 첫글자를 따서 소문자로 지음

// Deposit X amount on your account
func (a *Account) Deposit(amount int) {
	// (a Account)
	// GO가 여기서 a를 받아올 때 복사본을 만들어서 가져옴(safety를 위해, method가 속해있는 object를 보호하려고)
	// 즉, 이 프로젝트에서는 account의 복사본, 실제 account가 아니라는 소리
	// 복사본을 만들지 않고 그대로 account를 가져오고 싶으면 (a *Account): Deposit method를 호출한 account를 사용해라(pointer receiver)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// 에러 발생시키기 -> error.Error(){}
		// 아래와 같이 새로운 에러를 만들 수도 있음
		return errnoMoney
	}
	a.balance -= amount

	// error를 return 시켰으니 정상이어도 뭔가를 return해줘야함
	// error도 2가지 value가 있는데 하나는 error, 하나는 nil(null, None 같은 것)
	return nil
}

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string { // object가 너무 크면 복사본 만들지 말고 걍 pointer로 가져오렴
	return a.owner
}

func (a Account) String() string {
	// return "whatever you want"
	// Sprint는 string 출력 함수
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
