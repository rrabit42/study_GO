package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from home!") // 데이터를 Writer에 출력(Console에 출력 X)
}

func main() {
	// go로 서버 여는 방법 너무 쉬워요~
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // Fatal; os.Exit(1) 다음에 따라나오는 error를 Print() 하는 것과 동일, 에러가 있을 때만 실행될 것

}
