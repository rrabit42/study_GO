package main

import "fmt"

func main() {
	/*
		a := 2
		b := a
		a = 10 // b는 바뀌지 않음

		fmt.Println(a, b)
	*/

	a := 2
	b := &a // b는 a의 메모리 주소를 들여다 보고 있음

	fmt.Println(&a, b) // 동일함(a는 그냥 2로 출력됨)
	fmt.Println(*b)    // *는 loop through, see through의 의미! 즉, b가 들여다 보고 있는 a의 값인 2가 나옴

	a = 5
	fmt.Println(*b)

	*b = 20 // b를 이용해 a 값을 변화시킴
	fmt.Println(a)
}
