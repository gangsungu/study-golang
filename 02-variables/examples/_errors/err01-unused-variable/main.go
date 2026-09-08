// 일부러 컴파일에 실패시켜 에러 메시지를 확인하는 예제.
//
// 기대 에러:
//
//	declared and not used: a
//	declared and not used: msg
//	declared and not used: assigned
package main

func main() {
	var a int = 10
	var msg string = "Hello, Variable!"

	// 대입만 하는 것은 "사용"으로 치지 않는다. 값을 읽어야 한다.
	assigned := 1
	assigned = 2

	// 해결법 1) 실제로 사용한다  -> fmt.Println(a, msg, assigned)
	// 해결법 2) 빈 식별자로 버린다 -> _ = a
}
