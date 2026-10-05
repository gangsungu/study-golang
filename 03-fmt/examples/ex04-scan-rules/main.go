// Scan, Scanln, Scanf 가 입력을 어떻게 나누는지 비교한다.
//
// 표준 입력 대신 문자열을 읽는 Sscan 계열을 써서, 키보드 입력 없이 같은 결과를 재현한다.
// 규칙은 Scan 계열과 똑같다. (S = 문자열에서 읽음, 접두사 규칙은 출력과 동일)
package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("--- Scan: 개행도 공백으로 취급 ---")
	n, err := fmt.Sscan("1\n2", &a, &b)
	fmt.Println(n, err, a, b) // 2 <nil> 1 2

	fmt.Println("--- Scanln: 개행에서 멈춤 ---")
	a, b = 0, 0
	n, err = fmt.Sscanln("1\n2", &a, &b)
	fmt.Println(n, err, a, b) // 1 unexpected newline 1 0
	n, err = fmt.Sscanln("1 2 3", &a, &b)
	fmt.Println(n, err, a, b) // 2 expected newline 1 2  (남는 값이 있어도 에러)

	fmt.Println("--- Scanf: 서식과 정확히 일치해야 함 ---")
	n, err = fmt.Sscanf("1,2", "%d,%d", &a, &b)
	fmt.Println(n, err, a, b) // 2 <nil> 1 2
	n, err = fmt.Sscanf("1 2", "%d,%d", &a, &b)
	fmt.Println(n, err) // 1 input does not match format

	fmt.Println("--- 타입이 맞지 않으면 ---")
	n, err = fmt.Sscan("3 x", &a, &b)
	fmt.Println(n, err) // 1 expected integer  (앞의 값은 이미 저장됨)

	fmt.Println("--- 문자열은 공백에서 끊김 ---")
	var s string
	n, err = fmt.Sscan("hello world", &s)
	fmt.Println(n, err, s) // 1 <nil> hello

	fmt.Println("--- 포인터가 아니면 ---")
	n, err = fmt.Sscan("3", a) // 컴파일은 통과한다. 실행 중 에러로 돌아온다
	fmt.Println(n, err)        // 0 type not a pointer: int
}
