// 변수를 선언하는 여러 가지 방법을 확인한다.
package main

import "fmt"

// 패키지 레벨 변수는 := 로 선언할 수 없고 var를 써야 한다.
// 또한 사용하지 않아도 컴파일 에러가 나지 않는다.
var packageLevel = "함수 밖에서 선언한 변수"

func main() {
	// 1) 전체 형태
	var a int = 10

	// 2) 타입 생략 - 값에서 추론
	var b = 20

	// 3) 값 생략 - 제로값으로 초기화 (타입 필수)
	var c int

	// 4) 짧은 선언 - 함수 안에서만 가능
	d := 40

	// 5) 여러 개 한 번에. 타입이 달라도 된다.
	x, y := 1, "hello"

	// 6) 블록으로 묶기
	var (
		name string = "gopher"
		age  int    = 10
		ok   bool
	)

	fmt.Println(a, b, c, d)
	fmt.Println(x, y)
	fmt.Println(name, age, ok)
	fmt.Println(packageLevel)

	shortDeclRule()
	shadowing()
}

// := 는 왼쪽에 새 변수가 최소 하나 있어야 한다.
func shortDeclRule() {
	a, err := first()
	b, err := second() // OK. b가 새 변수라서 err는 대입으로 처리된다.
	// a, b := third() // 에러: no new variables on left side of :=

	fmt.Println(a, b, err)
}

// 안쪽 블록의 := 는 바깥 변수를 가리는 새 변수를 만든다 (섀도잉).
func shadowing() {
	v := "바깥"

	if true {
		v := "안쪽" // 바깥 v가 아니라 새로운 v
		fmt.Println("블록 안:", v)
	}

	fmt.Println("블록 밖:", v) // 여전히 "바깥"
}

func first() (int, error)  { return 1, nil }
func second() (int, error) { return 2, nil }
