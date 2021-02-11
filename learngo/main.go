package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string // array
	// method도 포함될 수 있음
}

/*
	python 의 __init__
	javascript의 constructor()
	Go는 이런 constructor method는 없음!
	struct는 constructor가 없기 때문에 우리 스스로 constructor를 실행해야 함. -> 나중에 연습할 것
*/

func main() {
	favFood := []string{"kimchi", "ramen"}
	// 이런 코드는 잘 쓰지 않음. 뭐가 뭔지 위에서 살펴봐야하기 때문
	// nico := person{"nico", 18, favFood}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
	fmt.Println(nico.age)
}
