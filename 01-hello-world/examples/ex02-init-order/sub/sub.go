// main 패키지가 아닌 평범한 라이브러리 패키지.
// import 되는 쪽이므로 main 패키지보다 먼저 초기화된다.
package sub

import "fmt"

// 패키지 변수의 초기화는 init()보다 먼저 일어난다.
var SubVar = trace("2. sub 패키지 변수 초기화")

func init() {
	fmt.Println("3. sub의 init()")
}

func trace(s string) string {
	fmt.Println(s)
	return s
}
