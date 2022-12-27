# study_GO  
~~졸업프로젝트를 위한 공부😤~~였다고 한다. 지금은 블록체인 업계에서 일하면서 필요함!  
* [노마드코더-GO 입문자](https://nomadcoders.co/go-for-beginners)  
* [한 눈에 끝내는 고랭 기초](https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88)  
* [노마드코인](https://nomadcoders.co/nomadcoin)  

> GO한테 반하는 중
> 1. 파이썬보다도 빠르고 간결함(multi-core processing, 병행성)  
> 2. return 값을 여러 개, 다른 타입으로 줄 수 있음  
> 3. naked return, if, switch에서 변수 선언 등 파이썬보다도 직관적인 언어인듯(쓸 수록 파이썬이 그리운걸 봐서 아닌듯ㅋㅋ)  
> 4. 그동안 python에 없어서 아쉬웠던 포인터까지 탑재...  

다만 정적 언어에 강타입 언어라는게 나를 간헐적으로 화나게 만든다... 그립구나 파이썬...  

## Go 프로그래밍 언어의 특성  
Go는 전통적인 컴파일, 링크 모델을 따르는 범용 프로그래밍 언어이다. Go는 일차적으로 시스템 프로그래밍을 위해 개발되었으며, C++, Java, Python의 장점들을 뽑아 만들어졌다. C++와 같이 Go는 컴파일러를 통해 컴파일되며, 정적 타입 (Statically Typed)의 언어이다. 또한 Java와 같이 Go는 Garbage Collection 기능을 제공한다. Go는 단순하고 간결한 프로그래밍 언어를 지향하였는데, Java의 절반에 해당하는 25개의 키워드만으로 프로그래밍이 가능하게 하였다. 마지막으로 Go의 큰 특징으로 Go는 Communicating Sequential Processes (CSP) 스타일의 Concurrent 프로그래밍을 지원한다.  


## Go 설치  
Go 프로그래밍을 시작하기 위해 Go 공식 웹사이트인 http://golang.org/dl 에서 해당 OS 버젼의 Go를 다운로드하여 설치한다. Go는 Windows, Linux, Mac OS X 에서 사용할 수 있다.  
윈도우즈에 Go를 설치하기 위해서는 MSI (*.msi) 파일을 다운받아 실행하면 되는데, Go는 디폴트로 ```C:\go``` 폴더에 설치되며, MSI가 ```C:\go\bin```을 PATH 환경변수를 추가한다.  
Go를 설치하고 해당 설치 디렉토리 밑에 bin 디렉토리를 보면 go.exe 파일이 있는데, 이 컴파일러로 go 프로그램을 컴파일하거나 실행할 수 있다. Go 프로그램은 파일 확장자 .go 를 갖는다.  


## Go 실행  
cmd에서 ```C:\GO\bin``` 디렉토리로 이동한 후 음과 같이 go 프로그램을 실행할 수 있다. Go는 run 명령어를 사용하면 직접 컴파일과 동시에 실행하게 된다. 이때 실행파일 .exe 는 만들어지지 않는다.  
> C:\Go\bin> go run test.go  

실행파일 .exe 을 생성하기 위해서는 다음과 같이 go build 명령을 사용한다.  

> C:\Go\bin> go build test.go  


## Workspace 폴더  
Go 프로그래밍을 위해 일반적으로 Workspace 폴더 (작업 폴더)를 생성하는데, 이 폴더 안에는 3개의 서브디렉토리 (bin, pkg, src)를 생성한다.  
예를 들어, C:\GoApp 디렉토리를 Workspace 폴더로 정했다면, C:\GoApp 안에 bin, pkg, src 서브 폴더를 만들어 준다.  

Workspace 폴더를 생성한 후, GOPATH 환경변수에 이 Workspace 폴더 경로를 추가해 준다.  
(SET GOPATH=C:\GoApp 처럼 세션별로 할 수 있으나, 주로 시스템 설정에서 시스템 환경변수 혹은 사용자 환경변수로 지정한다)  

GOPATH는 하나 이상의 경로를 지정할 수 있다. 즉, 여러 Workspace가 있는 경우, 이들 경로를 계속 추가할 수 있다.  

Go는 2개의 중요한 환경변수(GOROOT와 GOPATH)를 사용한다.  
* GOROOT: Go가 설치된 디렉토리(윈도우즈의 경우 디폴트로 C:\Go)를 가리키며, Go 실행파일은 GOROOT/bin 폴더에, Go 표준 패키지들은 GOROOT/pkg 폴더에 있다. (윈도우즈에 GO 설치시 시스템 환경변수로 자동으로 설정된다)  
* GOPATH: Go는 표준 패키지 이외의 3rd Party 패키지나 사용자 정의 패키지들을 이 GOPATH 에서 찾는다. 복수 개의 경로를 지정한 경우, 3rd Party 패키지는 처음 경로에 설치된다.  

Go 환경변수는 ```go env``` 를 실행하여 체크할 수 있다.  


## VS Code에서 Go 프로그래밍  
Visual Studio Code에서 Go Extension을 설치하면, Go 프로그래밍 편집기로 사용할 수 있다. 다음은 VS Code에서 Go를 사용하는 방법은 단계적으로 설명한 것이다.

1. VS Code 설치 (https://code.visualstudio.com/)

2. VS Code에 Go Extension을 설치한다. Extension 설치를 위해서는 VS Code에서 View -> Command Palette를 선택하고, ext install 이라고 치고 "Extensions: Install Extension"을 선택한다. 이후 계속 go 를 치고, 아래 화면처럼 lukehoban의 "Rich Go language support for VSC" 에서 다운로드 아이콘을 눌러 Extension을 설치한다. 설치 완료후 VS Code를 종료한다.

3. Go 작업폴더를 생성하고, GOPATH 환경변수에 이 폴더로 지정한다. Go 작업폴더는 일반적으로 bin, pkg, src 라는 3개의 서브디렉토리를 갖는다.  
```
GOPATH로 설정된 경로로 들어가서 bin, pkg, src 3개의 폴더를 생성한다.
- bin : go 에서 사용하는 명령어 저장
- pkg : go get 명령어로 다운 받은 패키지 저장
- src : go 파일 소스  
```  

4. VS Code를 실행한다. File -> Open Folder를 선택하고 Go 작업폴더를 연다. src 폴더에 새 test.go 파일을 만든다. go 파일이 만들어 지면, 화면하단에 아래와 같이 Analysis Tools Missing 링크가 표시되는데, 이를 클릭한다. 클릭하면, 상단에 Go analysis tools 설치 메시지가 뜨고 이를 설치한다.  
```
// main.go

package main
 
import (
    "fmt"
)
 
func main() {
    fmt.Println("hello world")
}
```  
작성해서 F5눌러서 디버그 콘솔에서 결과를 확인한다.  

5. 툴이 설치되면 Go 언어에 대해 인텔리센스와 키워드 컬러링이 제공된다.  

### 추가 설치  
아래의 명령어를 명령 프로픔트 창에 실행한다.  
To install the tools manually in the current GOPATH, just paste and run: [출처](https://github.com/Microsoft/vscode-go)  
```
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/newhook/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/cweill/gotests/...
go get -u -v golang.org/x/tools/cmd/godoc
go get -u -v github.com/fatih/gomodifytags
go get github.com/derekparker/delve/cmd/dlv // 얘는 go 디버거
```

### GO Modules  
GO 1.11 버전부터 생김.  
Go 모듈은 go 패키지들의 종속성을 관리하는 패키지 관리 시스템.  
Go의 모듈 개념은 Go 어플리케이션 내의 종속성 문제를 처리하기 위해 도입.  
모듈은 패키지를 트리 형식으로 저장하고 있으며, 루트 트리에는 go.mod 파일이 있다.  

저장소, 모듈 및 패키지 간의 관계는 아래와 같다.  
* 저장소에는 하나 이상의 Go 모듈이 포함됨.  
* 각 모듈에는 하나 이상의 Go 패키지가 포함되어 있음.  
* 각 패키지는 단일 디렉토리에 있는 하나 이상의 Go 소스 파일로 구성됨.  
```
repository
|-- module1
|   `-- package1
|       `-- src1.go
|       `-- src2.go
|   `-- go.mod
|   `-- package2
|       `-- src.go
|-- module2
|   `-- package
|       `--src.go
|   `-- go.mod
```

사용할 프로젝트 폴더에서 아래의 명령어를 쓰게 되면 go.mod란 파일이 생긴다  
```
go mod 프로젝트명
cd 프로젝트명
cat go.mod
# module github.com/내가작성한프로젝트명
```

### golang 모듈의 장점  
* $GOPATH/src 디렉토리 바깥에 프로젝트 디렉토리를 만들 수 있음
* vendor 디렉토리를 사용하지 않아도 됨
* reproducible build (언제,어디서나,누구라도 항상 동일한 build 결과를 보장함)

#### GO111MODULE 환경변수  
모듈 동작은 **GO111MODULE** 환경변수에 의해 제어된다.  
GO111MODULE은 on, off, auto(default)로 설정할 수 있다.  
* off: 빌드 중에 $GOPATH에 있는 패키지 사용  
* on: 빌드 중에 $GOPATH 대신 모듈에 있는 패키지 사용  
* auto: 현재 디렉터리가 $GOPATH 외부에 있고 go.mod 파일이 포함된 경우 모듈을 사용하고, 그렇지 않으면 $GOPATH 패키지를 사용  

아래의 명령어로 GO111MOUDLE 설정을 변경할 수 있다.  
```go env -w GO111MODULE=on```  

#### go.mod  
go.mod는 모듈을 정의하고 종속성 정보를 저장하고 있는 파일.  
이 파일을 통해 패키지들을 관리하게 된다.  

모듈은 루트 디렉터리에 하나의 go.mod 파일을 갖고 있다.  
go.mod는 네 가지 키워드 사용(module, require, replace, exclude)  
```
$ cat go.mod
module github.com/my/thing

go 1.15

require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency/v4 v4.0.0
)
```

* module  
module은 모듈 경로를 저장한다.  
위에 go.mod 파일 예를 보면 첫번째 라인에 모듈의 경로를 나타내는 module 지시자가 선언되어 있다.  
Go 소스 코드에서 패키지를 가져올 때, 절대 경로를 사용할 필요 없이 module에 선언되어 있는 모듈 경로를 사용하면 된다.  
```import "github.com/my/thing/bar"```  
예를 들어 모듈 안에 있는 bar 패키지를 가져온다면 위와 같이 선언해주면 된다.  

* require  
require 는 빌드 시 필요한 종속성 정보를 저장한다.
모듈을 사용하여 빌드하면 자동으로 필요한 패키지를 다운로드 및 설치하고 require에 패키지 경로와 버전 정보가 추가된다.

* replace  
replace는 모듈의 특정 버전이나, 버전 전체를 대체할 때 사용한다.  
require에 의해 설정된 종속성을, => 를 통해 우측에 설정된 패키지 버전으로 대체시킨다.  

빌드시 우측에 설정된 패키지 버전을 사용하게 된다.  
```replace example.com/some/dependency => example.com/some/dependency v1.2.3```  

* exclude  
exclude는 패키지의 특정 버전을 사용하지 않도록 할 때 사용한다.   

```exclude example.com/some/dependency```  

### go.sum
go.sum 파일은 go.mod에 종속성 정보가 추가될 때 생성.  
이 파일은 종속성 관리를 위한 암호화된 체크섬 정보를 저장하고 있고 각각의 체크섬을 확인하여 수정된 항목이 있는지 확인하는 데 사용.  
```
$ cat go.sum
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZO...
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:Nq...
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3...
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPX...
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/Q...
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9...
```

### module mirror  
모듈을 다운로드 하는 서버  
모듈의 소스레포가 죽었을 경우 (또는 폐기되었을 경우) 대비 (그런 경우에도 모듈 다운로드 보장)  
go 명령이 디폴트 참조하는 미러 서버를 구글이 운영 (proxy.golang.org)  
> 관련 명령 : go get, go build, ...  

### check sum DB  
모듈이 위/변조 되지 않았는지 확인하는 서버  
go 명령이 디폴트 참조하는 첵썸 디비를 구글이 운영 (sum.golang.org)  
관련 명령 : go get, gosumcheck, ...  

> go.mod 파일과 go.sum 파일은 소스관리 대상  

### 유용한 명령어  
* ```go mod init``` : 새로운 모듈 생성, go.mod 파일을 초기화
* ```go build, go test``` : 다른 패키지 빌딩 커맨드로써, 필요한 의존성을 go.mod에 추가한다.
* ```go list -m all``` : 현재 모듈의 의존성을 모두 출력
* ```go get``` : 특정 의존성의 필요한 버전으로 변경, 혹은 새로운 의존성 추가
* ```go mod tidy``` : 사용하지 않는 의존성 삭제

#### 참고  
https://ingeec.tistory.com/106  
https://yoongrammer.tistory.com/33  
https://soyoung-new-challenge.tistory.com/130  

