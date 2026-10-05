// 서식과 인자가 맞지 않는 Printf 를 go vet 이 잡아내는지 확인하는 예제.
//
// 컴파일 에러는 아니라서 go run 은 그대로 실행된다. 출력에 %!d(...) 같은 표시가 남을 뿐이다.
// 대신 go vet 이 실패한다. 이 폴더가 _errors/ 에 있는 이유다. (go vet ./... 를 깨뜨리지 않도록)
//
// 기대 에러 (go vet):
//
//	fmt.Printf format %d has arg name of wrong type string
//	fmt.Printf format %s reads arg #2, but call has 1 arg
//	fmt.Println call has possible Printf formatting directive %d
package main

import "fmt"

func main() {
	name := "gopher"

	fmt.Printf("이름: %d\n", name)         // 문자열에 %d
	fmt.Printf("이름: %s, 나이: %s\n", name) // 인자 부족
	fmt.Println("나이: %d", 10)            // Println 은 서식을 해석하지 않는다

	// 해결법) 서식 문자를 타입에 맞추고 개수를 맞춘다. 모르겠으면 %v
}
