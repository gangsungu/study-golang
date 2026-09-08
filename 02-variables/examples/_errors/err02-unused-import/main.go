// 일부러 컴파일에 실패시켜 에러 메시지를 확인하는 예제.
//
// 기대 에러:
//
//	"os" imported and not used
//	"strings" imported and not used
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("fmt만 쓰고 있다")

	// 해결법 1) 지운다
	// 해결법 2) 부수 효과만 필요하면 _ 를 붙인다 -> import _ "os"
}
