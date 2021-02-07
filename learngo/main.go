// 진입점: 컴파일러는 package 이름이 main인 것부터 찾음
// main은 컴파일을 위해서 필요함. 컴파일이 필요없다면 main.go 없어도 됨.
// 내가 작성할 pakage 이름
package main

import (
	"fmt" // fmt는 formmating을 위한 package

	"github.com/rrabit42/learngo/something"
)

func main() {
	fmt.Println("Hello world") // GO에서 function을 export하고 싶으면 대문자로 시작해야함
	something.SayHello()       // 다른 패키지로부터 exported function
	// something.sayBye() : private function
}
