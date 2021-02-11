package main

import "fmt"

func canIDrint(age int) bool {
	// GO에서는
	// 0. ()가 굳이 필요 없음
	// 1. if block은 return이 있으면 알아서 끝내줌. 즉 else가 필요 없음
	// 2. if 조건을 체크하기 전에 여기에 variable을 만들 수 있음(if에 사용하기 위해 변수를 만들었구나 한눈에 알 수 있음)
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
	// else {
	// 	return true
	// }
}

func main() {
	fmt.Println(canIDrint(16))
}
