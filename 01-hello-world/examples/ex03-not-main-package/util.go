// 패키지 이름이 main이 아니면 func main()이 있어도 진입점이 되지 않는다.
//
// 빌드는 통과한다:
//
//	go build ./01-hello-world/examples/ex03-not-main-package
//
// 하지만 실행하려 하면 거부된다:
//
//	go run ./01-hello-world/examples/ex03-not-main-package
//	-> package hello/01-hello-world/examples/ex03-not-main-package is not a main package
package util

import "fmt"

// 이름이 main일 뿐, 그냥 평범한 함수다.
// 소문자로 시작하므로 패키지 밖에서 호출할 수도 없다.
func main() {
	fmt.Println("이건 진입점이 아니다")
}

// Hello 는 대문자로 시작하므로 다른 패키지에서 호출할 수 있다.
func Hello() {
	fmt.Println("hi")
}
