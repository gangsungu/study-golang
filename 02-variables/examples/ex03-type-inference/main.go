// 리터럴이 어떤 타입으로 추론되는지 %T 로 확인한다.
package main

import "fmt"

func main() {
	a := 10
	b := 3.14
	c := 'A'
	d := "hi"
	e := true
	f := 3 + 4i

	fmt.Printf("10     -> %T\n", a)
	fmt.Printf("3.14   -> %T\n", b)
	fmt.Printf("'A'    -> %T (rune의 실체는 int32)\n", c)
	fmt.Printf("\"hi\"   -> %T\n", d)
	fmt.Printf("true   -> %T\n", e)
	fmt.Printf("3 + 4i -> %T\n", f)

	// 함정: 정수 리터럴은 float64가 아니라 int로 추론된다.
	g := 10
	h := 10.0
	fmt.Printf("\ng := 10   -> %T\n", g)
	fmt.Printf("h := 10.0 -> %T\n", h)

	// 'A' 는 문자가 아니라 숫자다. 그대로 출력하면 코드 포인트가 나온다.
	fmt.Printf("\nc 그대로 출력   : %v\n", c)
	fmt.Printf("string(c) 출력  : %v\n", string(c))
}
