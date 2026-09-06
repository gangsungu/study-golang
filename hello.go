package main

// 패키지 선언, 이 코드가 어떤 패키지에 속하는지 알려줌
// GoLang의 모든 코드는 반드시 패키지 선언으로 시작해야 함
// main() 패키지는 Go 프로그램의 진입점이 되는 패키지로, main 함수가 포함되어 있어야 함

import "fmt"

// 표준 입출력을 담당하는 내장 패키지

func main() {
	// main 함수는 Go 프로그램의 진입점으로, 프로그램 실행 시 가장 먼저 호출되는 함수
	fmt.Println("Hello, World!")
}
