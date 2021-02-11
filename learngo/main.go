package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0

	// range : array에 loop을 적용할 수 있도록 해줌(for 안에서만 사용 가능)
	for _, number := range numbers { // range는 index를 주기 때문에 이렇게 index를 첫번째로 받아줘야함(_ 하든지)
		total += number
	}
	return total

	// 당연히 이렇게도 가능!
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
