package main

import (
	"fmt"
	"time"
)

// Go routines란 기본적으로 다른 함수와 동시에 실행시키는 함수
//

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
	go sexyCount("nico")
	sexyCount("flynn")

	// 이렇게 해보면 쉽게 이해될 것.(프로그램이 다 실행되지 않아도 5초 지나면 프로그램 종료됨.)
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)

	// 메인함수와 go routine이 서로 정보를 주고 받으려면 어떻게 해야할까?
	// ex) url읆 체크하고 그 결과를 main에 보내야함.
	// main은 (주로) 정보를 저장하는 곳이니까!
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
