# study_GO  
졸업프로젝트를 위한 공부😤  

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

4. VS Code를 실행한다. File -> Open Folder를 선택하고 Go 작업폴더를 연다. src 폴더에 새 test.go 파일을 만든다. go 파일이 만들어 지면, 화면하단에 아래와 같이 Analysis Tools Missing 링크가 표시되는데, 이를 클릭한다. 클릭하면, 상단에 Go analysis tools 설치 메시지가 뜨고 이를 설치한다.

5. 툴이 설치되면 Go 언어에 대해 인텔리센스와 키워드 컬러링이 제공된다.
