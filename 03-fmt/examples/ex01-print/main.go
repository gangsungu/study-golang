// 출력 함수 세 가지(Print, Println, Printf)의 차이를 확인한다.
package main

import (
	"fmt"
	"os"
)

func main() {
	// Print: 피연산자를 그대로 이어 붙인다.
	// 공백은 "양쪽 모두 문자열이 아닐 때"만 넣는다. 개행은 없다.
	fmt.Print("a", "b", 1, 2, "c", 3.5, true)
	fmt.Print("\n")

	// Println: 피연산자 사이에 항상 공백을 넣고, 끝에 개행한다.
	fmt.Println("a", "b", 1, 2, "c", 3.5, true)

	// Printf: 서식 문자열대로 출력한다. 개행은 직접 넣어야 한다.
	name, age := "gopher", 10
	fmt.Printf("이름: %s, 나이: %d\n", name, age)

	// 세 함수 모두 (쓴 바이트 수, 에러)를 반환한다. 보통은 무시한다.
	n, err := fmt.Println("한글")
	fmt.Println("쓴 바이트:", n, "에러:", err) // "한글\n" = 3+3+1 = 7바이트

	// 접두사만 바꾸면 출력 대상이 달라진다.
	s := fmt.Sprintf("%s-%03d", "id", 7) // S: 문자열로 반환 (자바의 String.format)
	fmt.Println("Sprintf 결과:", s)
	fmt.Fprintln(os.Stderr, "Fprintln: 표준 에러로 출력") // F: 지정한 곳(io.Writer)으로 출력
}
