// 일부러 컴파일에 실패시켜 에러 메시지를 확인하는 예제.
//
// 기대 에러:
//
//	cannot use i (variable of type int) as float64 value in variable declaration
//	invalid operation: a + b (mismatched types int32 and int64)
//	invalid operation: n + m (mismatched types int and int64)
package main

import "fmt"

func main() {
	// 1) 범위가 넓어지는 방향이라도 암묵적 변환은 없다.
	var i int = 10
	var f float64 = i // float64(i) 필요

	// 2) 크기가 다른 정수끼리 연산 불가
	var a int32 = 1
	var b int64 = 2
	c := a + b // int64(a) + b 필요

	// 3) int와 int64는 크기가 같아도 다른 타입이다.
	var n int = 1
	var m int64 = 2
	d := n + m // int64(n) + m 필요

	fmt.Println(f, c, d)
}
