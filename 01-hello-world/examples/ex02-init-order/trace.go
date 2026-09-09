package main

import "fmt"

var mainVar = trace("4. main 패키지 변수 초기화")

// 같은 패키지 안에 init()이 여러 개 있어도 된다.
// 실행 순서는 파일 이름 순이라 main.go가 trace.go보다 먼저다.
// 다만 여기에 의존하는 코드를 쓰면 안 된다.
// 파일이 달라도 되고, 한 파일에 여러 개를 둬도 된다.
func init() {
	fmt.Println("6. trace.go의 init()")
}

func trace(s string) string {
	fmt.Println(s)
	return s
}
