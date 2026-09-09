package main

import "fmt"

// 같은 폴더 = 같은 패키지이므로 main을 두 번 선언할 수 없다.
func main() {
	fmt.Println("a.go의 main")
}
