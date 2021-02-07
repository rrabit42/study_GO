package main

import "fmt"

func main() {
	// type이 없는 constant(상수), 값 변경 불가
	// const name = "nico"
	// go는 type 언어, 무슨 타입인지 알려줘야함
	const name string = "nico"

	// variable
	// var name2 string = "nico"
	// name2 = "lynn" => 값 변경 가능
	name2 := "nico"
	// 같은 코드 축약형, type은 go가 알아서 찾아줌 그래서 type을 임의로 변경 못함
	// **첫번째 값의 type에 의존**해서 두번째 값의 type이 정해지기 때문
	// **func 안에서만 축약형 사용 가능. 변수에만 적용 가능**

	fmt.Println(name2)
}
