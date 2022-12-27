package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/rrabit42/study_GO/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string // 대소문자는 template에도 영향을 줌
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(w, "home", data)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func Start() {
	// go로 서버 여는 방법 너무 쉬워요~
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml")) // Template이나 error 반환하는 function의 호출 감싸서  error 있으면 에러 핸들링 해줌
	// templates이 모든 템플릿 가지고 있음. template과 templates 구분 주의
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // Fatal; os.Exit(1) 다음에 따라나오는 error를 Print() 하는 것과 동일, 에러가 있을 때만 실행될 것
}
