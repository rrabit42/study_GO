package main

import "fmt"

func canIDrint(age int) bool {
	// 익히 알다시피, 아래와 같이 쓸 수도 있지만
	// switch age {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// case 50:
	// 	return false
	// }

	// 이렇게 아예 대놓고 if-else 대체로도 쓸 수 있음
	// switch {
	// case age < 10:
	// 	return false
	// case age == 18:
	// 	return true
	// case age > 50:
	// 	return false
	// }
	// return false

	// if-else처럼 variable을 만들 수도 있음
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	case 50:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrint(18))
}
