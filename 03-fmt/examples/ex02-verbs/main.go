// 자주 쓰는 서식 문자(verb)와, 서식이 맞지 않을 때의 출력을 확인한다.
package main

import "fmt"

type point struct {
	X int
	Y string
}

func main() {
	p := point{1, "y"}

	// 무엇이든 출력하는 %v 계열
	fmt.Printf("%%v  : %v\n", p)  // 값만
	fmt.Printf("%%+v : %+v\n", p) // 필드 이름 포함
	fmt.Printf("%%#v : %#v\n", p) // Go 문법 그대로
	fmt.Printf("%%T  : %T\n", p)  // 타입

	// 정수
	fmt.Printf("%%d %%b %%o %%x %%X : %d %b %o %x %X\n", 255, 255, 255, 255, 255)
	fmt.Printf("%%c %%U        : %c %U\n", '한', '한') // 문자, 유니코드 코드포인트

	// 실수
	fmt.Printf("%%f %%e %%g    : %f %e %g\n", 1234.5678, 1234.5678, 1234.5678)

	// 문자열, bool, 포인터
	fmt.Printf("%%s %%q        : %s %q\n", "go", "go") // %q는 따옴표로 감싸고 이스케이프
	fmt.Printf("%%t           : %t\n", true)
	fmt.Printf("%%p           : %p\n", &p)
	fmt.Printf("%%%%           : %%\n") // % 문자 자체

	// 서식과 인자가 맞지 않아도 패닉이나 컴파일 에러가 나지 않는다.
	// 대신 출력 안에 표시가 남는다.
	// 인자를 슬라이스로 넘긴 것은 go vet 을 피하기 위해서다. 직접 쓰면 vet 이 잡아낸다 (_errors/err01 참고)
	wrongType := []any{"x"}
	tooFew := []any{1}
	tooMany := []any{1, 2}
	fmt.Printf("타입 불일치 : %d\n", wrongType...)  // %!d(string=x)
	fmt.Printf("인자 부족   : %d %d\n", tooFew...) // 1 %!d(MISSING)
	fmt.Printf("인자 초과   : %d", tooMany...)     // 1%!(EXTRA int=2)
	fmt.Println()
}
