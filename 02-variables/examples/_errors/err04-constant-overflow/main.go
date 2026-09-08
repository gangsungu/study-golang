// 일부러 컴파일에 실패시켜 에러 메시지를 확인하는 예제.
//
// 상수는 컴파일러가 값을 알고 있으므로, 손실이 생기는 변환을 미리 막는다.
// 같은 연산이라도 변수로 하면 그냥 통과한다는 점이 핵심이다.
//
// 기대 에러:
//
//	cannot convert 3.9 (untyped float constant) to type int
//	cannot use 300 (untyped int constant) as int8 value in variable declaration (overflows)
package main

import "fmt"

func main() {
	// 1) 상수의 손실 변환은 컴파일 타임에 막힌다.
	truncated := int(3.9)

	// 2) 상수 오버플로우도 마찬가지다.
	var small int8 = 300

	// 반면 아래는 통과한다. 런타임 값이라 컴파일러가 판단할 수 없기 때문이다.
	//   var f float64 = 3.9
	//   fmt.Println(int(f))     // 3
	//   var big int32 = 300
	//   fmt.Println(int8(big))  // 44

	fmt.Println(truncated, small)
}
