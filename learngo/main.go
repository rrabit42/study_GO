package main

import (
	"fmt"
	"strings"
)

// return 값의 type을 {} 앞에다가 요롷게. 이 type이 없으면 go는 return값으로 아무것도 주지 않는다고 생각.
// func multiply(a int, b int) int {}
// 인자들이 모두 같은 타입이면 이렇게도 가능
func multiply(a, b int) int {
	return a * b
}

// go의 func은 multiple return이 가능하다.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 원하는 만큼의 arguments를 전달받는 법
func repeatMe(words ...string) {
	fmt.Println(words) // array 형태로 출력
}

func main() {
	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("nico") // 참고로 go는 무언가를 만들고 쓰지 않으면 에러 발생

	// return 개수에 맞게 변수 줘야함. if not 에러
	// totalLen := lenAndUpper("sy") => X
	totalLen, _ := lenAndUpper("sy") // 이런식으로 작성하면 value 값을 무시하는 것(ignored value). 컴파일러가 쳐다보지도 않고 무시

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	fmt.Println(totalLength, upperName)
	fmt.Println(totalLen)
}
