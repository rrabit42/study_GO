package accounts

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

// Go의 패턴
// constructor가 없기 때문에 function으로 construct하거나 struct를 만들도록 한다.
// obejct을 만들어서 return 시키는 함수

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // object 자체를 return, 복사하는걸 원하지 않아서!
}
