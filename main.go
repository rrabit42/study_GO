package main

import (
	"fmt"
	"time"
)

// RULE OF CHANNELS & GOROUTINE
// 1. 메인함수가 끝이나면 고루틴은 무용지물
// 2. 채널을 통해 어떤 타입의 데이터를 주고 받을 건지 지정해줘야함.
// 3. <- : 데이터를 받는 다는 뜻, BLOCKING OPERATION! (참고로 데이터를 받을 곳이 없어도 채널을 통해 데이터를 보낼 수는 있음.)

func main() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i)
		// go routine이 완료가 되지 않았으면 기다림; blocking operation
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		fmt.Println(person)
		time.Sleep(time.Second)
	}
}
