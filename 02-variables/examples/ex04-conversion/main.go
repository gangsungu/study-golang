// 타입 변환이 어떤 경우에 필요하고, 값이 어떻게 조용히 바뀌는지 확인한다.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1) 범위가 넓어지는 방향이라도 명시적 변환이 필요하다.
	var i int = 10
	var f float64 = float64(i)
	fmt.Println("int -> float64:", f)

	// 2) 크기가 다르면 같은 정수끼리도 바로 연산할 수 없다.
	var a int32 = 1
	var b int64 = 2
	// c := a + b        // 에러: mismatched types int32 and int64
	c := int64(a) + b
	fmt.Println("int32 + int64:", c)

	// 3) int와 int64는 크기가 같아도 다른 타입이다.
	var n int = 5
	var n64 int64 = int64(n)
	fmt.Println("int -> int64:", n64)

	// 4) 실수 -> 정수는 반올림이 아니라 절삭이다.
	//    단 int(3.9) 처럼 상수 리터럴을 직접 변환하는 것은 컴파일 에러다.
	//    상수 변환은 값이 정확히 표현될 때만 허용되기 때문이다.
	var pos float64 = 3.9
	var neg float64 = -3.9
	fmt.Println("\nint(3.9)  =", int(pos), "(반올림이면 4여야 한다)")
	fmt.Println("int(-3.9) =", int(neg), "(0 방향으로 잘린다)")

	// 5) 범위를 벗어나면 비트가 잘려나가지만 에러는 나지 않는다.
	//    이것도 변수라서 통과한다. var x int8 = 300 은 컴파일 에러다.
	var big int32 = 300
	fmt.Println("int8(300) =", int8(big), "(300 - 256 = 44)")

	// 6) 숫자 <-> 문자열은 변환이 아니라 strconv를 쓴다.
	fmt.Println("\nstring(rune(65)) =", strconv.Quote(string(rune(65))), "<- 코드 포인트로 해석된다")
	fmt.Println("strconv.Itoa(65) =", strconv.Quote(strconv.Itoa(65)), "<- 보통 원하는 것")

	num, err := strconv.Atoi("65")
	fmt.Println("strconv.Atoi(\"65\") =", num, err)

	_, err = strconv.Atoi("육십오")
	fmt.Println("잘못된 입력:", err)

	fmt.Println("FormatFloat:", strconv.FormatFloat(3.14159, 'f', 2, 64))
}
