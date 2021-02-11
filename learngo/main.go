package main

import (
	"fmt"
	"strings"
)

// naked return : return할 variable을 굳이 꼭 명시하지 않아도 됨
// return type과 함께 변수명도 같이 정의해 주는 것!
// 무엇이 리턴되는지 정확히 알 수 있음
func lenAndUpper(name string) (length int, uppercase string) {
	// return 후에 작업을 시키고 싶을 때는
	// defer를 쓴다. 'defer'은 function 값이 return하고 나면 실행됨
	defer fmt.Println("I'm done")

	length = len(name) // update, not create 때문에 :=를 쓰는게 아님
	uppercase = strings.ToUpper(name)
	return
	// return length, uppercase 이렇게 해도 당연히 작동하지만, 굳이 작성하지 않아도 됨
}

func main() {
	totalLength, up := lenAndUpper("nico") // 즉 여기 function이 끝나면 defer가 실행됨
	fmt.Println(totalLength, up)           // 그 후에 출력!
}
