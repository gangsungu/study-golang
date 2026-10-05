// 입력을 한 줄 단위로 읽은 뒤 파싱한다.
// 실패해도 그 줄만 버려지므로 다음 입력에 찌꺼기가 남지 않는다.
//
// ex05 와 같은 입력("3 a b", "5 6")으로 비교해 본다.
// (또는 printf '3 a b\n5 6\n' | go run ./03-fmt/examples/ex06-read-line)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin) // 입력 스트림을 줄 단위로 끊어 주는 도구

	for i := 1; i <= 2; i++ {
		fmt.Printf("%d번째 입력 (정수 2개): ", i)
		if !sc.Scan() { // 한 줄을 읽는다. 더 읽을 것이 없으면 false
			fmt.Println("  -> 입력 끝")
			return
		}
		line := sc.Text() // 개행(\n, \r\n)은 제거된 상태

		var a, b int
		n, err := fmt.Sscanln(line, &a, &b) // 읽어 둔 줄을 파싱
		if err != nil {
			fmt.Printf("  -> 에러: %v | 읽은 개수: %d | 버려진 줄: %q\n", err, n, line)
			continue
		}
		fmt.Println("  -> 성공:", a, b)
	}
}
