// 잘못된 입력이 표준 입력 스트림에 남아서 다음 입력까지 망가뜨리는 것을 확인한다.
//
// 실행 후 첫 줄에 "3 a b" 를, 두 번째 줄에 "5 6" 을 입력해 본다.
// (또는 printf '3 a b\n5 6\n' | go run ./03-fmt/examples/ex05-stdin-leftover)
package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		var a, b int
		fmt.Printf("%d번째 입력 (정수 2개): ", i)
		n, err := fmt.Scanln(&a, &b)
		if err != nil {
			fmt.Println("  -> 에러:", err, "| 읽은 개수:", n)
			continue
		}
		fmt.Println("  -> 성공:", a, b)
	}
	// "3 a b" 를 넣으면
	//   1번째: 3은 읽고 a에서 실패. a는 소비되어 사라지고 " b\n" 이 스트림에 남음
	//   2번째: 남아 있던 b를 읽다가 실패 -> 사용자는 입력한 적도 없는데 에러
	//   3번째: 남아 있던 개행만 읽고 실패 (unexpected newline)
	//   4번째: 그제야 "5 6" 을 읽음
}
