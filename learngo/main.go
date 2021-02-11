package main

import "fmt"

func main() {
	// map은 key와 value로 이루어진 데이터
	// [key] 는 string이고 value{} 도 string이다
	nico := map[string]string{
		"name": "nico",
		// "age": 12 -> string, string이라고 선언했기에 쓸 수 없음 => object와 다른 부분
		"age": "12",
	}
	fmt.Println(nico)

	// map도 range를 통해 다룰 수 있다.
	for key, value := range nico {
		fmt.Println(key, value)
	}

	// map에 데이터 추가, 검색 등 여러 함수가 있지만 coming soon
}
