## Go 시작하기

### 모든 Go 파일은 패키지 선언으로 시작함
+ 파일 첫 줄은 반드시 `package 이름` 이어야 함
    - 주석은 그 위에 올 수 있지만, 코드로서는 패키지 선언이 가장 먼저임
    - 패키지 = 코드를 묶는 단위이며, **폴더 하나가 패키지 하나**임
+ 같은 폴더 안의 모든 `.go` 파일은 패키지 이름이 같아야 함
    - 한 폴더에 `package main` 과 `package util` 을 섞으면 컴파일 에러
+ 패키지 이름은 관례상 소문자 한 단어를 씀
    - 언더스코어나 카멜케이스를 쓰지 않음 (`httputil` O, `http_util` X)

### main 패키지와 main 함수
+ 실행 가능한 프로그램이 되려면 **두 조건을 모두** 만족해야 함
    - 패키지 이름이 `main`
    - 그 안에 매개변수도 반환값도 없는 `func main()` 이 존재
+ 둘 중 하나라도 빠지면 실행 파일이 만들어지지 않음
    - `package util` 에 `func main()` 을 써도 에러는 안 나지만, 그냥 이름이 main인 평범한 함수일 뿐임
    - 라이브러리 패키지로 취급되어 `go run` 이 동작하지 않음
+ `main()` 은 프로그램 진입점이며 인자를 받지 않음
    - C나 Java와 달리 매개변수로 명령행 인자를 받지 않음
    - 명령행 인자는 `os.Args` 로 읽음

### import
+ 다른 패키지를 가져다 쓸 때 사용함
    - `fmt` 는 표준 입출력을 담당하는 내장 패키지
    - `Println`, `Printf` 처럼 **대문자로 시작하는 이름만** 외부에서 접근 가능함 (Go의 접근 제어 방식)
+ 여러 개를 가져올 때는 괄호로 묶음

    ```go
    import (
        "fmt"
        "os"
        "strings"
    )
    ```

+ **사용하지 않는 import는 컴파일 에러임**
    - `imported and not used: "os"`
    - 변수와 같은 규칙이며, 자세한 내용은 [변수 선언과 사용](../02-variables/변수-선언과-사용.md) 참고
+ 부수 효과(init 함수 실행)만 필요하면 `_` 를 붙여 이름 없이 가져옴

    ```go
    import _ "github.com/lib/pq"
    ```

### 세미콜론과 중괄호
+ 문장 끝에 세미콜론을 쓰지 않음
    - 실제로는 컴파일러가 줄 끝에 자동으로 넣어줌 (세미콜론 자동 삽입)
+ 그래서 **여는 중괄호를 다음 줄에 쓸 수 없음**

    ```go
    func main()
    {           // 에러: func main() 뒤에 세미콜론이 자동 삽입되어 버림
    }
    ```

    - K&R 스타일(같은 줄에 `{`)이 문법적으로 강제되는 셈임
+ 포맷은 `gofmt` 가 강제함
    - 들여쓰기는 스페이스가 아니라 **탭**
    - 취향 논쟁이 없도록 표준 포맷터 하나로 통일한 것이 Go의 방침
    - 대부분의 에디터가 저장 시 자동 실행되도록 설정되어 있음

### 실행 방법
+ `go run` — 컴파일 후 바로 실행하고 실행 파일은 남기지 않음
    - 학습·확인용

    ```bash
    go run ./01-hello-world/examples/ex01-hello
    go run main.go              # 파일을 직접 지정해도 됨
    ```

+ `go build` — 실행 파일을 만듦
    - 결과물 이름은 기본적으로 **폴더 이름**을 따라감 (`tutorial/` 에서 빌드하면 `tutorial.exe`)
    - `-o` 로 이름을 지정할 수 있음

    ```bash
    go build -o hello.exe ./01-hello-world/examples/ex01-hello
    ```

+ `go vet` — 컴파일은 되지만 의심스러운 코드를 잡아줌
+ `go fmt` — 소스를 표준 포맷으로 정렬함

### go.mod
+ 모듈(프로젝트) 단위를 정의하는 파일이며, 레포 루트에 하나 둠

    ```
    module hello

    go 1.27.1
    ```

+ `module` 은 이 프로젝트의 이름이자 import 경로의 접두사임
    - `hello/02-variables/examples/ex01-declaration` 같은 형태로 조합됨
+ **import 경로에는 한글을 쓸 수 없음**
    - `malformed import path: invalid char '변'`
    - 그래서 이 레포는 `.go` 파일이 들어가는 폴더 이름을 영문으로 둠
+ `go mod init 모듈이름` 으로 생성함

### 정리
+ 파일 구조
    - 패키지 선언 → import → 코드 순서
+ 실행되려면
    - `package main` + `func main()` 둘 다 필요
+ 컴파일 에러가 되는 것
    - 사용하지 않는 import, 사용하지 않는 지역 변수
    - 여는 중괄호를 다음 줄에 쓰는 것
+ 자주 쓰는 명령
    - `go run` 실행, `go build` 빌드, `go vet` 검사, `go fmt` 포맷

### 관련 문서
+ [변수 선언과 사용](../02-variables/변수-선언과-사용.md) — 미사용 변수·import 규칙
