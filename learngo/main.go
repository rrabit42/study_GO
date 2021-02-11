package main

import "fmt"

func main() {
	// array
	// names := [5]string{"nico", "lynn", "dal"}
	// names[3] = "alala"
	// fmt.Println(names)
	// 이 상태로도 출력은 됨.

	// slice : 데이터 타입 중 하나, 기본적으로 array 인데 length 없이 사용하는 것
	// 요소 추가, 삭제가 자유로움. 추가할 때 array처럼 index로 지정이 불가하니 append(slice, value) 이용
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn") // append는 modify가 아님. 새로운 값이 추가된 slice를 return함
	fmt.Println(names)
}
