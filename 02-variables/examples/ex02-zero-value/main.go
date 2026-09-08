// 초기화하지 않은 변수가 어떤 값을 갖는지 확인한다.
package main

import "fmt"

type point struct {
	X int
	Y string
}

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var p *int
	var sl []int
	var m map[string]int
	var arr [3]int
	var st point

	fmt.Printf("int     : %d\n", i)
	fmt.Printf("float64 : %v\n", f)
	fmt.Printf("bool    : %t\n", b)
	fmt.Printf("string  : %q (빈 문자열이지 nil이 아니다)\n", s)
	fmt.Printf("*int    : %v\n", p)
	fmt.Printf("[]int   : %v (nil == %t)\n", sl, sl == nil)
	fmt.Printf("map     : %v (nil == %t)\n", m, m == nil)
	fmt.Printf("[3]int  : %v (요소마다 제로값)\n", arr)
	fmt.Printf("struct  : %+v (필드마다 제로값)\n", st)

	// nil 슬라이스는 그대로 append 할 수 있다.
	sl = append(sl, 1)
	fmt.Println("\nnil 슬라이스에 append:", sl)

	// nil 맵은 읽기는 되지만 쓰면 런타임 패닉이 난다.
	fmt.Println("nil 맵 읽기:", m["없는키"], "(제로값이 나온다)")
	// m["키"] = 1 // panic: assignment to entry in nil map
	m = make(map[string]int)
	m["키"] = 1
	fmt.Println("make 후 쓰기:", m)
}
