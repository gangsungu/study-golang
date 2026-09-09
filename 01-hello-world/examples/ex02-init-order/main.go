// 초기화 순서를 확인한다.
//
//	import한 패키지 변수 -> 그 패키지의 init() -> 내 패키지 변수 -> 내 init() -> main()
//
// 실행: go run ./01-hello-world/examples/ex02-init-order
package main

import (
	"fmt"

	// 이름을 쓰지 않고 init()만 실행시키기 위한 blank import.
	// 드라이버 등록에 흔히 쓰이는 방식이다.
	_ "hello/01-hello-world/examples/ex02-init-order/sub"
)

func init() {
	fmt.Println("5. main.go의 init()")
}

func main() {
	fmt.Println("7. main() - 여기서 비로소 진입점이 실행된다")
	fmt.Println("\nmainVar =", mainVar)
}
