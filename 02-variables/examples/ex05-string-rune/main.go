// 문자열이 UTF-8 바이트 시퀀스라는 사실에서 파생되는 동작들을 확인한다.
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "한글abc"

	// 1) len()은 문자 수가 아니라 바이트 수다.
	fmt.Println("문자열      :", s)
	fmt.Println("len()       :", len(s), "<- 바이트 수 (한글 1자 = 3바이트)")
	fmt.Println("[]rune 길이 :", len([]rune(s)), "<- 문자 수")
	fmt.Println("RuneCount   :", utf8.RuneCountInString(s))

	// 2) 인덱싱하면 문자가 아니라 byte가 나온다.
	fmt.Printf("\ns[0]        : %v (타입 %T)\n", s[0], s[0])
	fmt.Printf("string(s[0]): %q <- 한글 1바이트만 잘려 깨진다\n", string(s[0]))

	// 3) range는 rune 단위로 돌고, 인덱스는 바이트 위치다.
	fmt.Println("\nrange 순회:")
	for i, r := range s {
		fmt.Printf("  인덱스 %d -> %q (코드 포인트 %d)\n", i, r, r)
	}

	// 4) 문자열은 불변이다. 바꾸려면 []rune 이나 []byte 로 변환해야 한다.
	// s[0] = 'x' // 에러: cannot assign to s[0]
	r := []rune(s)
	r[0] = '韓'
	fmt.Println("\n[]rune 으로 수정:", string(r))

	// 5) 반복 연결은 strings.Builder 를 쓴다.
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		sb.WriteString("go")
	}
	fmt.Println("strings.Builder :", sb.String())
}
