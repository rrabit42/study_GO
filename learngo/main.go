// 진입점: 컴파일러는 package 이름이 main인 것부터 찾음
// main은 컴파일을 위해서 필요함. 컴파일이 필요없다면 main.go 없어도 됨.
// 내가 작성할 pakage 이름
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
