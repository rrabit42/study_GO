package main

import (
	"fmt"
	"time"
)

// Go routines란 기본적으로 다른 함수와 동시에 실행시키는 함수

func main() {

	// Top-Down 방법: 순차적으로 실행
	// sexyCount("nico")
	// sexyCount("flynn")

	// 병렬적으로 실행: go만 붙여주면 됨 -> nico func를 flynn func과 나란히 놓은 셈, go만 더해줌으로서 작업들이 동시에 진행됨.

	// go를 flynn 앞에도 추가하면 아무일도 일어나지 않고 프로그램이 종료됨, 왜냐하면 메인 함수가 작업을 마쳤기 때문
	// goroutine은 프로그램이 작동하는 동안만 유효함. 다시 말해 메인 함수가 실행되는 동안만!
	// 이 경우에는 첫번째 go routine 실행하고, 두번째 go routine 실행하고 끝나게 된 것.(메인 함수는 할 일이 없음. 모두 go routine)
	// 메인 함수는 go routine을 기다려주지 않음!!!!!! 메인 함수가 끝나면 go routine도 죽음
	// 아래 예제가 되는 경우는 메인 함수가 flynn을 카운팅하고 있기 때문.
	// go sexyCount("nico")
	// sexyCount("flynn")

	// 이렇게 해보면 쉽게 이해될 것.(프로그램이 다 실행되지 않아도 5초 지나면 프로그램 종료됨.)
	// go sexyCount("nico")
	// go sexyCount("flynn")
	// time.Sleep(time.Second * 5)

	// 메인함수와 go routine이 서로 정보를 주고 받으려면 어떻게 해야할까?
	// ex) url읆 체크하고 그 결과를 main에 보내야함.
	// main은 (주로) 정보를 저장하는 곳이니까!

	// Channel은 (goroutine-메인함수), (goroutine-다른 goroutine) 사이에 정보를 전달하기 위한 방법 -> go routine 사이 어떻게 커뮤니케이션할까에 대한 해답
	c := make(chan string) // 어떤 종류의 정보를 주고 받을건지 정의, main 함수 안에 있으니 main과 소통하기 위한 채널이 되겠다.
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		// 값을 리턴하고 있다고 해서 아래와 같이 사용할 수는 없음
		// result := go isSexy(person)
		go isSexy(person, c)
	}
	fmt.Println("Waiting for messages")
	// resultOne := <-c    // time sleep이 없어도 main 함수가 기다림: 채널로부터 무언가를 받을 때 메인함수는 어떤 답이 올 때까지 기다림
	// resultTwo := <-c
	// resultThree := <-c	// go는 컴파일 전에는 이걸 알아낼 수 없음. blocking operation이기 때문에 런타임에서 기다리면서 알아냄.
	// 프로그램에는 남아 있는 goroutine이 없는데 하염없이 메세지를 기다린다면 프로그램은 안끝나고 뭐가 잘못됐는지도 모르는 지경이 될 수 있음.
	// 그리고 개수만큼 계속 변수 할당에서 이렇게 쓸거야?!
	// -> loop를 쓰자!

	// 메인 함수는 아래 함수를 보자 멈출 것. 왜냐하면 이건 blocking operation 이기 때문, 이 작업이 끝날 때 까지 멈춘다는 뜻
	fmt.Println("Received this message:", <-c) // 데이터 하나만 나옴, channel에는 2개 존재
	// 다른 메세지 받을 때까지 또 기다림
	fmt.Println("Received this message:", <-c) // 이렇게 쓸 수도 있다. routine 개수에 맞지 않게 channel에서 데이터 가져오려 하면 Deadlock 에러 남

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy" // return true; go routine에 return을 받는 대신 channel을 통해 메세지를 보내는 것
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		fmt.Println(person)
		time.Sleep(time.Second)
	}
}
